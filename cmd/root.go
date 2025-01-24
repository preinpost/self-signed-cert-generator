package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var test string

var rootCmd = &cobra.Command{
	Use: "cert",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("RUN")
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&test, "test-name", "t", "", "common usage")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error Occured %s", err)
		os.Exit(1)
	}
}
