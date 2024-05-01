package vm

import (
	"fmt"
	"os"
	"proxcli/colors"
	"proxcli/config"
	"proxcli/filter"
	"proxcli/request"
	"strconv"

	"github.com/spf13/cobra"
)

func Vmstart(id int, name string) {
	config := config.InitConfig()
	switch {
	case name != "none":
		exist := filter.Exist(name, group, id)
		if !exist {
			fmt.Printf("\u274C ERROR: No Vm with name %s found\n", colors.Red(name))
		} else {
			id = filter.GetId(name)
			vm := filter.Vminfo(id)
			if vm.Data.Status == "running" {
				fmt.Printf("\u2705 %s Vm is already running\n", vm.Data.Name)
			} else {
				url := fmt.Sprintf("https://%s:8006/api2/json/nodes/%s/qemu/%s/status/start", config["ip"], config["node"], strconv.Itoa(id))
				data, code := request.NewRequest(url, "POST")
				if code == 200 {
					fmt.Printf("\u2705 %s Vm started\n", vm.Data.Name)
				} else {
					fmt.Printf("\u274C %d - ERROR %s\n", code, data)
				}
			}
		}
	default:
		exist := filter.Exist(name, group, id)
		if !exist {
			fmt.Printf("\u274C ERROR: No Vm with id %s found\n", colors.Red(strconv.Itoa(id)))
		} else {
			vm := filter.Vminfo(id)
			if vm.Data.Vmid == id {
				if vm.Data.Status == "running" {
					fmt.Printf("\u2705 %s Vm is already running\n", vm.Data.Name)
				} else {
					url := fmt.Sprintf("https://%s:8006/api2/json/nodes/%s/qemu/%s/status/start", config["ip"], config["node"], strconv.Itoa(id))
					data, code := request.NewRequest(url, "POST")
					if code == 200 {
						fmt.Printf("\u2705 %s Vm started\n", vm.Data.Name)
					} else {
						fmt.Printf("\u274C %d - ERROR %s\n", code, data)
					}
				}
			}
		}
	}

}

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
		Vmstart(id, name)
	},
}
