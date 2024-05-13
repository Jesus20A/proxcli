/*
Copyright © 2024 Jesus Salas
*/
package vm

import (
	"proxcli/cmd"

	"github.com/spf13/cobra"
)

var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "Interact with VMs",
	Long:  `Perform actions on VMs`,
}

var state string
var group string
var id int
var name string

func init() {
	cmd.RootCmd.AddCommand(vmCmd)
	vmCmd.AddCommand(VmsList)
	vmCmd.AddCommand(VmGet)
	vmCmd.AddCommand(VmStart)
	vmCmd.AddCommand(VmStop)
	vmCmd.AddCommand(Group)
	Group.AddCommand(GetGroup)
	Group.AddCommand(ListGroup)
	Group.AddCommand(StartGroup)
	Group.AddCommand(StopGroup)
	VmsList.Flags().StringVarP(&state, "state", "s", "all", "Vm state (running, stopped, all)")
	VmGet.Flags().IntVarP(&id, "id", "i", -1, "Vm id")
	VmGet.Flags().StringVarP(&name, "name", "n", "none", "Vm name (Only if inventory configured)")
	VmStart.Flags().StringVarP(&name, "name", "n", "none", "Vm name (Only if inventory configured)")
	VmStart.Flags().IntVarP(&id, "id", "i", -1, "Vm id")
	VmStop.Flags().StringVarP(&name, "name", "n", "none", "Vm name (Only if inventory configured)")
	VmStop.Flags().IntVarP(&id, "id", "i", -1, "Vm id")
	GetGroup.Flags().StringVarP(&group, "name", "n", "none", "Group name (Only if inventory configured)")
	StartGroup.Flags().StringVarP(&group, "name", "n", "none", "Group name (Only if inventory configured)")
	StopGroup.Flags().StringVarP(&group, "name", "n", "none", "Group name (Only if inventory configured)")

}
