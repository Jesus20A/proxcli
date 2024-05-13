/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package inventory

import (
	"fmt"
	"log"
	"proxcli/cmd"
	"proxcli/cmd/lxc"
	"proxcli/cmd/vm"
	"proxcli/pkg/config"
	"proxcli/pkg/filter"
	"strconv"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func Inventory(source string) {

	switch source {
	case "vms":
		v := vm.Vmsinfo("false", "silent")

		out, err := yaml.Marshal(v)
		if err != nil {
			log.Fatal(err)
		} else if err := filter.WritetoFile(string(out), config.Inventoryfile); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("\u2705 VMs added to inventory file at %s\n", config.Inventoryfile)
		}
	case "lxc":
		l := lxc.Lxcsinfo("false", "silent")
		container := &l
		for i := range len(container.Lxc) {
			container.Lxc[i].Id, _ = strconv.Atoi(container.Lxc[i].Id.(string))
		}

		out, err := yaml.Marshal(l)
		if err != nil {
			log.Fatal(err)
		} else if err := filter.WritetoFile(string(out), config.Inventoryfile); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("\u2705 Lxc containers added to inventory file at %s\n", config.Inventoryfile)
		}

	case "all":
		v := vm.Vmsinfo("false", "silent")

		out, err := yaml.Marshal(v)
		if err != nil {
			log.Fatal(err)
		} else if err := filter.WritetoFile(string(out), config.Inventoryfile); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("\u2705 VMs added to inventory file at %s\n", config.Inventoryfile)
		}

		l := lxc.Lxcsinfo("false", "silent")
		container := &l
		for i := range len(container.Lxc) {
			container.Lxc[i].Id, _ = strconv.Atoi(container.Lxc[i].Id.(string))
		}

		out, err = yaml.Marshal(l)
		if err != nil {
			log.Fatal(err)
		} else if err := filter.WritetoFile(string(out), config.Inventoryfile); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("\u2705 Lxc containers added to inventory file at %s\n", config.Inventoryfile)
		}
	default:
		fmt.Println("Invalid source")
	}

}

// inventoryCmd represents the inventory command
var inventoryCmd = &cobra.Command{
	Use:   "inventory",
	Short: "Create inventory file",
	Long:  `Create inventory file of VMs and Lxc containers`,
	Run: func(cmd *cobra.Command, args []string) {
		Inventory(source)
	},
}

var source string

func init() {
	cmd.RootCmd.AddCommand(inventoryCmd)
	inventoryCmd.Flags().StringVarP(&source, "source", "s", "all", "Source of inventory: vms, lxc, all")
}
