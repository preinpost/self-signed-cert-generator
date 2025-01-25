package cmd

import (
	"cert-demo/pkg/intermediate"
	"cert-demo/pkg/rootca"
	"cert-demo/pkg/server"
	"cert-demo/pkg/utils"

	"github.com/spf13/cobra"
)

// ROOT Flags
var RootOrganizationName string

// Intermediate Flags
var InterOrganizationName string

// Server Flags
var CommonName string

var ServerOrganizationName string
var San []string

var newCmd = &cobra.Command{
	Use:           "new",
	Short:         "새로운 인증서 세트 생성",
	Long:          "rootCA, intermediateCA, server 그리고 chaining 인증서를 생성 합니다.",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		// if CommonName == "" {
		// 	return errors.New("'--common-name'이 설정되지 않았습니다")
		// }

		rootca.GenRootCert(RootOrganizationName)
		intermediate.GenIntermidiateCert(InterOrganizationName)
		server.GenServerCert(ServerOrganizationName, San)

		utils.ChainingCert()

		return nil
	},
}

func init() {
	newCmd.Flags().StringVar(&RootOrganizationName, "root-organization-name", "", "인증서 RootCA Organization Name")
	newCmd.Flags().StringVar(&InterOrganizationName, "inter-organization-name", "", "인증서 IntermediateCA Organization Name")
	// newCmd.Flags().StringVar(&CommonName, "common-name", "", "인증서 Common Name")

	newCmd.Flags().StringVar(&ServerOrganizationName, "server-organization-name", "", "인증서 Server Organization Name")
	newCmd.Flags().StringArrayVar(&San, "san", []string{}, "서버 인증서 Subject Alternative Name")

	newCmd.MarkFlagRequired("root-organization-name")
	newCmd.MarkFlagRequired("inter-organization-name")

	// newCmd.MarkFlagRequired("common-name")
	newCmd.MarkFlagRequired("server-organization-name")
	newCmd.MarkFlagRequired("san")

	rootCmd.AddCommand(newCmd)
}
