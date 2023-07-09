package add

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add is a palette that contains commands to add on to a vault",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to gopwd!")
	},
}

func Execute() {
	err := addCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
