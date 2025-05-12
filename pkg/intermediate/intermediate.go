package intermediate

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"

	"github.com/preinpost/self-signed-cert-generator/pkg/utils"
)

type IntermediateCertParams struct {
	Organization string
}

var IntermediateCert *x509.Certificate
var IntermediateKey *ecdsa.PrivateKey

func GenIntermidiateCert(certPath, keyPath string, params *IntermediateCertParams) {

	caCert, err := utils.LoadCert(certPath)
	if err != nil {
		panic("ca cert read error")
	}

	caKey, err := utils.LoadKey(keyPath)
	if err != nil {
		panic("ca key read error")
	}

	intermediateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	intermediateTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject: pkix.Name{
			Organization: []string{params.Organization},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(90, 0, 0), // 90년 유효
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLen:            1,
	}

	intermediateCertDER, _ := x509.CreateCertificate(rand.Reader, intermediateTemplate, caCert, &intermediateKey.PublicKey, caKey)
	intermediateCert, _ := x509.ParseCertificate(intermediateCertDER)

	IntermediateCert = intermediateCert
	IntermediateKey = intermediateKey

	utils.SaveCertAndKey("intermediate.pem", "intermediate.key", intermediateCertDER, intermediateKey)
}
