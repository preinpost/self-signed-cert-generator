package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "cert",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {

		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		fmt.Fprintf(os.Stderr, "see command help")
		os.Exit(1)
	}
}
