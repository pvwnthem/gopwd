package vault

import (
	"os"

	"github.com/pvwnthem/gopwd/pkg/util"
	"github.com/spf13/cobra"
)

var configFile string

var (
	Path  string
	Force bool
)

var VaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "Vault is a palette that contains commands to manage and create vaults",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	Path, configFile, _ = util.InitConfig(Path, configFile)

	err := VaultCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	VaultCmd.PersistentFlags().StringVarP(&Path, "path", "p", "", "The path to create the vault at")
	VaultCmd.PersistentFlags().BoolVarP(&Force, "yes", "y", false, "if the command asks for confirmation, force it to yes")

}
