package vault

import (
	"fmt"

	"github.com/spf13/cobra"
)

var insertCmd = &cobra.Command{
	Use:   "insert [site]",
	Short: "Inserts a new password into the vault",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		site := args[0]
		fmt.Println("Inserting password for", site)
	},
}

func init() {
	VaultCmd.AddCommand(insertCmd)
}
