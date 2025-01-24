package main

import "cert-demo/cmd"

var Outdir = "./out"

func main() {
	// rootca.GenRootCert()
	// intermideate.GenIntermidiateCert()
	// server.GenServerCert()

	// utils.ChainingCert()

	cmd.Execute()
}
