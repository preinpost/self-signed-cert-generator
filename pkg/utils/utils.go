package utils

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var CertOutdir = "./out"

func SaveCertAndKey(certFileName, keyFileName string, certDER []byte, key crypto.Signer) {
	if err := os.Mkdir(CertOutdir, os.ModePerm); err != nil && !os.IsExist(err) {
		fmt.Println("오류 발생:", err)
	}

	certOut, _ := os.Create(fmt.Sprintf("%s/%s", CertOutdir, certFileName))
	defer certOut.Close()
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	keyOut, _ := os.Create(fmt.Sprintf("%s/%s", CertOutdir, keyFileName))
	defer keyOut.Close()

	bytes, err := x509.MarshalECPrivateKey(key.(*ecdsa.PrivateKey))

	if err != nil {
		panic(err)
	}

	pem.Encode(keyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: bytes})
}

func ChainingCert() {
	serverCert, err := os.ReadFile(fmt.Sprintf("%s/%s", CertOutdir, "server.pem"))
	if err != nil {
		log.Fatalf("Failed to read server certificate: %v", err)
	}

	// Intermediate CA 인증서 읽기
	intermediateCert, err := os.ReadFile(fmt.Sprintf("%s/%s", CertOutdir, "intermediate.pem"))
	if err != nil {
		log.Fatalf("Failed to read intermediate certificate: %v", err)
	}

	// 두 인증서를 결합
	combinedCert := append(serverCert, intermediateCert...)

	// 결합된 인증서를 파일로 저장
	err = os.WriteFile(fmt.Sprintf("%s/%s", CertOutdir, "chaining.pem"), combinedCert, 0644)
	if err != nil {
		log.Fatalf("Failed to write combined certificate: %v", err)
	}

	log.Println("server.pem과 intermediate.pem 으로 chaining 인증서를 생성하였습니다.")
}

func LoadCert(certPath string) (*x509.Certificate, error) {
	dat, err := os.ReadFile(filepath.Join(CertOutdir, certPath))

	if err != nil {
		return nil, fmt.Errorf("인증서 파일 읽기를 실패했습니다.")
	}

	// PEM 디코딩
	block, _ := pem.Decode(dat)
	if block == nil {
		return nil, fmt.Errorf("PEM 디코딩 실패: 유효한 PEM 데이터가 없습니다")
	}

	// x509.Certificate로 파싱
	crt, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("인증서 파싱 실패: %w", err)
	}

	return crt, nil
}

func LoadKey(filePath string) (interface{}, error) {
	// 파일 경로 처리 (플랫폼 호환성 보장)
	fullPath := filepath.Join(CertOutdir, filePath)

	// 파일 읽기
	pemBytes, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("PrivateKey 파일 읽기 실패: %w", err)
	}

	// PEM 디코딩
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("PEM 디코딩 실패: 유효한 PEM 데이터가 없습니다")
	}

	// 키 형식에 따라 파싱
	switch block.Type {
	case "EC PRIVATE KEY": // ECDSA 키
		key, err := x509.ParseECPrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("ECDSA 키 파싱 실패: %w", err)
		}
		return key, nil

	case "RSA PRIVATE KEY": // RSA 키
		key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("RSA 키 파싱 실패: %w", err)
		}
		return key, nil

	case "PRIVATE KEY": // PKCS#8 형식의 키 (RSA 또는 ECDSA)
		key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("PKCS#8 키 파싱 실패: %w", err)
		}
		return key, nil

	default:
		return nil, fmt.Errorf("지원하지 않는 키 형식: %s", block.Type)
	}
}
