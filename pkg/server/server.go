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

func GenServerCert(organization string, san []string) {
	serverKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	serverTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(3),
		Subject: pkix.Name{
			Organization: []string{organization},
			Country:      []string{"KR"},
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(50, 0, 0), // 50년 유효
		KeyUsage:    x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},

		DNSNames: san,
	}
	serverCertDER, _ := x509.CreateCertificate(rand.Reader, serverTemplate, intermideate.IntermediateCert, &serverKey.PublicKey, intermideate.IntermediateKey)
	utils.SaveCertAndKey("server.pem", "server.key", serverCertDER, serverKey)
}
