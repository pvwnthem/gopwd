package vault

import (
	"os"

	"github.com/pvwnthem/gopwd/util"
	"github.com/spf13/cobra"
)

var configFile string

var (
	Path string
	Name string
)

var VaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "Vault is a palette that contains commands to manage and create vaults",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	err := VaultCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() { Path, Name, configFile, _ = util.InitConfig(Path, Name, configFile) })
	VaultCmd.PersistentFlags().StringVarP(&Path, "path", "p", "", "The path to create the vault at")
	VaultCmd.PersistentFlags().StringVarP(&Name, "name", "n", "", "The name of the vault")

}
