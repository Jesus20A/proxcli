package lxc

import (
	"os"
	"proxcli/pkg/stop"

	"github.com/spf13/cobra"
)

var LxcStop = &cobra.Command{
	Use:   "stop",
	Short: "Stop Container",
	Long:  `Stop a Container by specifying its ID or, if the inventory is configured, also by Name`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		if id == -1 && name == "none" {
			cmd.Help()
			os.Exit(0)
		}
		stop.Stop(id, name, "lxc")
	},
}
