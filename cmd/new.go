package cmd

import (
	"cert-demo/pkg/intermediate"
	"cert-demo/pkg/rootca"
	"cert-demo/pkg/server"
	"cert-demo/pkg/utils"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "새로운 인증서 세트 생성",
	Long:  "rootCA, intermediateCA, server 그리고 chaining 인증서를 생성 합니다.",
	Run: func(cmd *cobra.Command, args []string) {
		rootca.GenRootCert()
		intermediate.GenIntermidiateCert()
		server.GenServerCert()

		utils.ChainingCert()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
