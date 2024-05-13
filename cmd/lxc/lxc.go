package lxc

import (
	"proxcli/cmd"

	"github.com/spf13/cobra"
)

// lxcCmd represents the lxc command
var lxcCmd = &cobra.Command{
	Use:   "lxc",
	Short: "Interact with Lxc containers",
	Long:  `Perform actions on Lxc containers`,
}
var state string
var id int
var name string

func init() {
	cmd.RootCmd.AddCommand(lxcCmd)
	lxcCmd.AddCommand(LxcList)
	lxcCmd.AddCommand(LxcGet)
	lxcCmd.AddCommand(LxcStart)
	lxcCmd.AddCommand(LxcStop)
	LxcList.Flags().StringVarP(&state, "state", "s", "all", "Container state (running, stopped, all)")
	LxcGet.Flags().IntVarP(&id, "id", "i", -1, "Container id")
	LxcGet.Flags().StringVarP(&name, "name", "n", "none", "Container name")
	LxcStart.Flags().StringVarP(&name, "name", "n", "none", "Container name")
	LxcStart.Flags().IntVarP(&id, "id", "i", -1, "Container id")
	LxcStop.Flags().StringVarP(&name, "name", "n", "none", "Container name")
	LxcStop.Flags().IntVarP(&id, "id", "i", -1, "Container id")
}
