package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gopwd",
	Short: "A cli password manager written in go",
	Long:  "gopwd is an encrypted cli password manager (similar to password-store) written in golang",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to gopwd!")
	},
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
