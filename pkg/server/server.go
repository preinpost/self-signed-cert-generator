package server

import (
	intermideate "cert-demo/pkg/intermediate"
	"cert-demo/pkg/utils"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"
)

func GenServerCert() {
	serverKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	serverTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(3),
		Subject: pkix.Name{
			Organization: []string{"Server"},
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(1, 0, 0), // 1년 유효
		KeyUsage:    x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},

		DNSNames: []string{"localhost"},
	}
	serverCertDER, _ := x509.CreateCertificate(rand.Reader, serverTemplate, intermideate.IntermediateCert, &serverKey.PublicKey, intermideate.IntermediateKey)
	utils.SaveCertAndKey("server.pem", "server.key", serverCertDER, serverKey)
}
