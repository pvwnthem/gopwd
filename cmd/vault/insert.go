package vault

import (
	"github.com/spf13/cobra"
)

var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "inserts a new password into the vault",
	Long:  "",

	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

func init() {
	VaultCmd.AddCommand(insertCmd)
}
