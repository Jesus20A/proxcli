package vm

import (
	"fmt"
	"os"
	"proxcli/colors"
	"proxcli/config"
	"proxcli/filter"
	"proxcli/request"
	"proxcli/structs"
	"strconv"
	"strings"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func groupStart(group string) {
	config := config.InitConfig()
	exist := filter.Exist(name, group, id)
	if !exist {
		fmt.Printf("\u274C ERROR: No Group with name %s found\n", colors.Red(group))
	} else {
		vms := filter.Getgroup(group)
		for _, v := range vms {
			exist := filter.Exist("none", "false", v.Id)
			if !exist {
				fmt.Printf("\u274C ERROR: No Vm with id %s found\n", colors.Red(strconv.Itoa(v.Id)))
			} else if vm := filter.Vminfo(v.Id); vm.Data.Name == v.Name {
				if vm.Data.Status == "running" {
					fmt.Printf("\u2705 %s Vm is already running\n", vm.Data.Name)
				} else {
					url := fmt.Sprintf("https://%s:8006/api2/json/nodes/%s/qemu/%s/status/start", config["ip"], config["node"], strconv.Itoa(v.Id))
					data, code := request.NewRequest(url, "POST")
					if code == 200 {
						fmt.Printf("\u2705 %s Vm started\n", vm.Data.Name)
					} else {
						fmt.Printf("\u274C %d - ERROR %s\n", code, data)
					}
				}
			} else {
				fmt.Printf("\u274C Error: Vm name %s and id %s do not match\n", colors.Red(v.Name), colors.Red(strconv.Itoa(v.Id)))
			}
		}

	}
}

func groupStop(group string) {
	config := config.InitConfig()
	exist := filter.Exist(name, group, id)
	if !exist {
		fmt.Printf("\u274C ERROR: No Group with name %s found\n", colors.Red(group))
	} else {
		vms := filter.Getgroup(group)
		for _, v := range vms {
			exist := filter.Exist("none", "false", v.Id)
			if !exist {
				fmt.Printf("\u274C ERROR: No Vm with id %s found\n", colors.Red(strconv.Itoa(v.Id)))
			} else if vm := filter.Vminfo(v.Id); vm.Data.Name == v.Name {
				if vm.Data.Status == "stopped" {
					fmt.Printf("\u2705 %s Vm is already stopped\n", vm.Data.Name)
				} else {
					url := fmt.Sprintf("https://%s:8006/api2/json/nodes/%s/qemu/%s/status/shutdown", config["ip"], config["node"], strconv.Itoa(v.Id))
					data, code := request.NewRequest(url, "POST")
					if code == 200 {
						fmt.Printf("\u2705 %s Vm stopped\n", vm.Data.Name)
					} else {
						fmt.Printf("\u274C %d - ERROR %s\n", code, data)
					}
				}
			} else {
				fmt.Printf("\u274C Error: Vm name %s and id %s do not match\n", colors.Red(v.Name), colors.Red(strconv.Itoa(v.Id)))
			}
		}

	}

}

func getGroup(group string) {
	exist := filter.Exist(name, group, id)
	if !exist {
		fmt.Printf("\u274C ERROR: No Group with name %s found\n", colors.Red(group))
	} else {
		vms := filter.Getgroup(group)
		data := []structs.VmInfo{}
		for _, v := range vms {
			exist := filter.Exist("none", "false", v.Id)
			if !exist {
				fmt.Printf("\u274C No Vm with id %s found\n", colors.Red(strconv.Itoa(v.Id)))
			} else if vm := filter.Vminfo(v.Id); vm.Data.Name == v.Name {
				info := filter.Vminfo(v.Id)
				data = append(data, info)
			} else {
				fmt.Printf("\u274C Error: Vm name %s and id %s do not match\n", colors.Red(v.Name), colors.Red(strconv.Itoa(v.Id)))
				os.Exit(1)
			}
		}
		Table(data)

	}
}

func listGroup() {
	data, err := os.ReadFile(config.Inventoryfile)
	if err != nil {
		fmt.Println(err)
	}
	groups := structs.Groups{}
	if err := yaml.Unmarshal(data, &groups); err != nil {
		fmt.Println(err)
	}
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Name"},
			{Align: simpletable.AlignCenter, Text: "Vms"},
		},
	}

	var cells [][]*simpletable.Cell

	var vms []string
	for _, group := range groups.Groups {

		for _, g := range group.Vms {
			vms = append(vms, g.Name)
		}
		cells = append(cells, []*simpletable.Cell{
			{Text: colors.White(group.Name)},
			{Text: colors.Blue(strings.Join(vms, ", "))},
		})
		table.Body = &simpletable.Body{Cells: cells}
		table.SetStyle(simpletable.StyleRounded)
		vms = vms[:0]
	}
	table.Println()
}

var GetGroup = &cobra.Command{
	Use:   "get",
	Short: "Get group members info",
	Long:  `Display more detailed information about the VMs, which are part of the group`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if isSet := cmd.Flags().Lookup("name").Changed; !isSet {
			cmd.Help()
			os.Exit(0)
		}
		getGroup(name)
	},
}

var StartGroup = &cobra.Command{
	Use:   "start",
	Short: "Start VMs Group",
	Long:  `Start the VMs in a group, specifying the group name`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if isSet := cmd.Flags().Lookup("name").Changed; !isSet {
			cmd.Help()
			os.Exit(0)
		}
		groupStart(name)
	},
}

var StopGroup = &cobra.Command{
	Use:   "stop",
	Short: "Stop VMs Group",
	Long:  `Stop the VMs in a group, specifying the group name`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if isSet := cmd.Flags().Lookup("name").Changed; !isSet {
			cmd.Help()
			os.Exit(0)
		}
		groupStop(name)
	},
}

var ListGroup = &cobra.Command{
	Use:   "list",
	Short: "List all the groups",
	Long:  `List the groups and their members`,
	Run: func(cmd *cobra.Command, args []string) {
		listGroup()
	},
}

var Group = &cobra.Command{
	Use:   "group",
	Short: "Interact whit groups",
	Long:  `perform actions on a group if it is configured in the inventory file`,
}
