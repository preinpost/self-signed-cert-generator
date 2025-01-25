package server

import (
	"cert-demo/pkg/utils"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"
)

type ServerCertParams struct {
	Organization string
	San          []string
}

func GenServerCert(certPath, keyPath string, params *ServerCertParams) {

	caCert, err := utils.LoadCert(certPath)
	if err != nil {
		panic("ca cert read error")
	}

	caKey, err := utils.LoadKey(keyPath)
	if err != nil {
		panic("ca key read error")
	}

	serverKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	serverTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(3),
		Subject: pkix.Name{
			Organization: []string{params.Organization},
			Country:      []string{"KR"},
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(50, 0, 0), // 50년 유효
		KeyUsage:    x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},

		DNSNames: params.San,
	}
	serverCertDER, _ := x509.CreateCertificate(rand.Reader, serverTemplate, caCert, &serverKey.PublicKey, caKey)
	utils.SaveCertAndKey("server.pem", "server.key", serverCertDER, serverKey)
}
