package configure

import (
	"proxcli/cmd"
	"proxcli/pkg/config"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Create configuration file",
	Long:  `Create configuration directory if not present and ask the information needed to connect to the proxmox node`,
	Run: func(cmd *cobra.Command, args []string) {
		config.CreateConfig()
	},
}

func init() {
	cmd.RootCmd.AddCommand(configureCmd)
}
