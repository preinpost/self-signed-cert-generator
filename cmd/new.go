package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var fileName string

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "short explain",
	Long:  "Long explain",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("filename = %s\n", fileName)
	},
}

func init() {
	newCmd.Flags().StringVarP(&fileName, "filename", "f", "", "attach file")
	rootCmd.AddCommand(newCmd)
}
