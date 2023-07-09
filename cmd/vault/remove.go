package vault

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pvwnthem/gopwd/util"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes an existing vault at the specified path",
	Long:  "",

	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if the vault exists
		if _, err := os.Stat(filepath.Join(Path, Name)); os.IsNotExist(err) {
			return errors.New("vault does not exist")
		}

		err := os.RemoveAll(filepath.Join(Path, Name))
		if err != nil {
			return fmt.Errorf("failed to remove vault: %w", err)
		}

		return nil
	},
}

func init() {
	// Add flags for the remove command (similar to the init command)
	removeCmd.Flags().StringVarP(&Path, "path", "p", filepath.Join(util.GetHomeDir(), ".gopwd"), "The path to the vault")
	removeCmd.Flags().StringVarP(&Name, "name", "n", "", "The name of the vault")

	VaultCmd.AddCommand(removeCmd)

	err := removeCmd.MarkFlagRequired("name")
	if err != nil {
		panic(err)
	}
}
