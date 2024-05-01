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

func Vmstop(id int, name string) {
	config := config.InitConfig()
	switch {
	case name != "none":
		exist := filter.Exist(name, group, id)
		if !exist {
			fmt.Printf("\u274C ERROR: No Vm with name %s found\n", colors.Red(name))
		} else {
			id = filter.GetId(name)
			vm := filter.Vminfo(id)
			if vm.Data.Status == "stopped" {
				fmt.Printf("\u2705 %s Vm is already stopped\n", vm.Data.Name)
			} else {
				url := fmt.Sprintf("https://%s:8006/api2/json/nodes/%s/qemu/%s/status/shutdown", config["ip"], config["node"], strconv.Itoa(id))
				data, code := request.NewRequest(url, "POST")
				if code == 200 {
					fmt.Printf("\u2705 %s Vm stopped\n", vm.Data.Name)
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
			if vm.Data.Status == "stopped" {
				fmt.Printf("\u2705 %s Vm is already stopped\n", vm.Data.Name)
			} else {
				url := fmt.Sprintf("https://%s:8006/api2/json/nodes/%s/qemu/%s/status/shutdown", config["ip"], config["node"], strconv.Itoa(id))
				data, code := request.NewRequest(url, "POST")
				if code == 200 {
					fmt.Printf("\u2705 %s Vm stopped\n", vm.Data.Name)
				} else {
					fmt.Printf("\u274C %d - ERROR %s\n", code, data)
				}

			}

		}
	}
}

var VmStop = &cobra.Command{
	Use:   "stop",
	Short: "Stop Vm",
	Long:  `Stop a Vm by specifying its ID or, if the inventory is configured, also by Name`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		if id == -1 && name == "none" {
			cmd.Help()
			os.Exit(0)
		}
		Vmstop(id, name)
	},
}
