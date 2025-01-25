package cmd

import (
	"github.com/spf13/cobra"
)

var caFileName string

var generateCmd = &cobra.Command{
	Use:   "genserver",
	Short: "server 용 인증서 생성",
	Long:  "RootCa 또는 Intermediate CA를 사용해서 server 인증서를 생성합니다.",
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

func init() {
	generateCmd.Flags().StringVar(&caFileName, "rootca", "", "attach file")

	rootCmd.AddCommand(generateCmd)
}
