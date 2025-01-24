package main

import (
	intermideate "cert-demo/pkg/intermediate"
	"cert-demo/pkg/rootca"
	"cert-demo/pkg/server"
	"cert-demo/pkg/utils"
)

var Outdir = "./out"

func main() {
	rootca.GenRootCert()
	intermideate.GenIntermidiateCert()
	server.GenServerCert()

	utils.ChainingCert()
}
