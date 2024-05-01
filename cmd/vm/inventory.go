package vm

import (
	"fmt"
	"proxcli/config"
	"proxcli/filter"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func inventory() {
	c := Vmsinfo("false", "silent")
	out, err := yaml.Marshal(c)
	if err != nil {
		cobra.CheckErr(err)
	} else if err := filter.WritetoFile(string(out), config.Inventoryfile); err != nil {
		cobra.CheckErr(err)
	} else {
		fmt.Printf("\u2705 Inventory file created at %s\n", config.Inventoryfile)
	}

}

var Inventory = &cobra.Command{
	Use:   "inventory",
	Short: "Create inventory file",
	Long:  `Create inventory file with all currently available VMs`,
	Run: func(cmd *cobra.Command, args []string) {
		inventory()
	},
}
