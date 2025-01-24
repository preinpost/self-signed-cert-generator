package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CommonName string

var rootCmd = &cobra.Command{
	Use: "cert",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if CommonName == "" {
			return fmt.Errorf("'--common-name'이 설정되지 않았습니다")
		}
		return nil
	}

	rootCmd.PersistentFlags().StringVarP(&CommonName, "common-name", "c", "", "인증서 Common Name")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error Occured %s", err)
		os.Exit(1)
	}
}
