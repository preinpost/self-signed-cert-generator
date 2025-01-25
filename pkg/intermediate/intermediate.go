package intermediate

import (
	"cert-demo/pkg/rootca"
	"cert-demo/pkg/utils"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"
)

var IntermediateCert *x509.Certificate
var IntermediateKey *ecdsa.PrivateKey

func GenIntermidiateCert(organization string) {
	intermediateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	intermediateTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject: pkix.Name{
			Organization: []string{organization},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(90, 0, 0), // 90년 유효
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLen:            1,
	}

	intermediateCertDER, _ := x509.CreateCertificate(rand.Reader, intermediateTemplate, rootca.RootCert, &intermediateKey.PublicKey, rootca.RootKey)
	intermediateCert, _ := x509.ParseCertificate(intermediateCertDER)

	IntermediateCert = intermediateCert
	IntermediateKey = intermediateKey

	utils.SaveCertAndKey("intermediate.pem", "intermediate.key", intermediateCertDER, intermediateKey)
}
