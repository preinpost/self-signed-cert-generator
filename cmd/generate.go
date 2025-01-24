package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var intermediateCaFileName string

var generateCmd = &cobra.Command{
	Use:   "genserver",
	Short: "server 용 인증서 생성",
	Long:  "RootCa 또는 Intermediate CA를 사용해서 server 인증서를 생성합니다.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("intermediateCaFileName = %s\n", intermediateCaFileName)
	},
}

func init() {
	generateCmd.Flags().StringVarP(&intermediateCaFileName, "rootca", "r", "", "attach file")
	rootCmd.AddCommand(generateCmd)
}
