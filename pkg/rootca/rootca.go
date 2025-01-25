package rootca

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

var RootCert *x509.Certificate
var RootKey *ecdsa.PrivateKey

func GenRootCert(organization string) {
	rootKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	rootTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{organization},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(100, 0, 0), // 100년 유효
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLen:            2,
	}

	rootCertDER, _ := x509.CreateCertificate(rand.Reader, rootTemplate, rootTemplate, &rootKey.PublicKey, rootKey)
	rootCertValue, _ := x509.ParseCertificate(rootCertDER)

	RootCert = rootCertValue
	RootKey = rootKey
	utils.SaveCertAndKey("root.pem", "root.key", rootCertDER, rootKey)
}
