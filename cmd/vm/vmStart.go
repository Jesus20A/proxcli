package vm

import (
	"os"
	"proxcli/pkg/start"

	"github.com/spf13/cobra"
)

var VmStart = &cobra.Command{
	Use:   "start",
	Short: "Start Vm",
	Long:  `Start a Vm by specifying its ID or, if the inventory is configured, also by Name`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		if id == -1 && name == "none" {
			cmd.Help()
			os.Exit(0)
		}
		start.Start(id, name, "qemu")
	},
}
