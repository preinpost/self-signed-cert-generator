package utils

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
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
