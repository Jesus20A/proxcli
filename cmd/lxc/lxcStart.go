package lxc

import (
	"os"
	"proxcli/pkg/start"

	"github.com/spf13/cobra"
)

var LxcStart = &cobra.Command{
	Use:   "start",
	Short: "Start container",
	Long:  `Start a Container by specifying its ID or, if the inventory is configured, also by Name`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		if id == -1 && name == "none" {
			cmd.Help()
			os.Exit(0)
		}
		start.Start(id, name, "lxc")
	},
}
