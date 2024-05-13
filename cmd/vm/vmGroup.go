package vm

import (
	"fmt"
	"log"
	"os"
	"proxcli/pkg/colors"
	"proxcli/pkg/config"
	"proxcli/pkg/filter"
	"proxcli/pkg/start"
	"proxcli/pkg/stop"
	"proxcli/pkg/types"
	"strconv"
	"strings"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func groupStart(group string) {
	vms, err := filter.Getgroup(group)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range vms {
		vm, err := filter.Vminfo(v.Id)
		if err != nil {
			log.Fatal(err)
		}
		if vm.Data.Name == v.Name {
			if vm.Data.Status == "running" {
				fmt.Printf("\u2705 %s Vm is already running\n", vm.Data.Name)
			} else {
				start.Start(v.Id, "none", "qemu")
			}
		} else {
			fmt.Printf("\u274C Error: Vm name %s and id %s do not match\n", colors.Red(v.Name), colors.Red(strconv.Itoa(v.Id)))
		}
	}

}

func groupStop(group string) {
	vms, err := filter.Getgroup(group)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range vms {
		vm, err := filter.Vminfo(v.Id)
		if err != nil {
			log.Fatal(err)
		}
		if vm.Data.Name == v.Name {
			if vm.Data.Status == "stopped" {
				fmt.Printf("\u2705 %s Vm is already stopped\n", vm.Data.Name)
			} else {
				stop.Stop(v.Id, "none", "qemu")
			}
		} else {
			log.Fatalf("\u274C Error: Vm name %s and id %s do not match\n", colors.Red(v.Name), colors.Red(strconv.Itoa(v.Id)))

		}
	}

}

func grouptable(data []types.VmInfo) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "id"},
			{Align: simpletable.AlignCenter, Text: "Name"},
			{Align: simpletable.AlignCenter, Text: "Cpu"},
			{Align: simpletable.AlignCenter, Text: "Mem"},
			{Align: simpletable.AlignCenter, Text: "Status"},
		},
	}

	var cells [][]*simpletable.Cell

	for _, item := range data {

		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", item.Data.Vmid)},
			{Text: colors.Blue(item.Data.Name)},
			{Text: colors.White(strconv.Itoa(item.Data.Cpus))},
			{Text: colors.Yellow(filter.MemConverter(float32(item.Data.Memory)))},
			{Text: colors.Color(item.Data.Status, item.Data.Status)},
		})

	}

	table.Body = &simpletable.Body{Cells: cells}
	table.SetStyle(simpletable.StyleRounded)
	table.Println()

}
func getGroup(group string) {
	vms, err := filter.Getgroup(group)
	if err != nil {
		log.Fatal(err)
	}
	data := []types.VmInfo{}
	for _, v := range vms {
		vm, err := filter.Vminfo(v.Id)
		if err != nil {
			log.Fatal(err)
		}
		if vm.Data.Name == v.Name {
			data = append(data, vm)
		} else {
			fmt.Printf("\u274C Error: Vm name %s and id %s do not match\n", colors.Red(v.Name), colors.Red(strconv.Itoa(v.Id)))
			os.Exit(1)
		}
	}
	grouptable(data)

}

func listGroup() {
	data, err := os.ReadFile(config.Inventoryfile)
	if err != nil {
		fmt.Println(err)
	}
	groups := types.Groups{}
	if err := yaml.Unmarshal(data, &groups); err != nil {
		log.Fatal(err)
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
		group, _ := cmd.Flags().GetString("name")
		if isSet := cmd.Flags().Lookup("name").Changed; !isSet {
			cmd.Help()
			os.Exit(0)
		}
		getGroup(group)
	},
}

var StartGroup = &cobra.Command{
	Use:   "start",
	Short: "Start VMs Group",
	Long:  `Start the VMs in a group, specifying the group name`,
	Run: func(cmd *cobra.Command, args []string) {
		group, _ := cmd.Flags().GetString("name")
		if isSet := cmd.Flags().Lookup("name").Changed; !isSet {
			cmd.Help()
			os.Exit(0)
		}
		groupStart(group)
	},
}

var StopGroup = &cobra.Command{
	Use:   "stop",
	Short: "Stop VMs Group",
	Long:  `Stop the VMs in a group, specifying the group name`,
	Run: func(cmd *cobra.Command, args []string) {
		group, _ := cmd.Flags().GetString("name")
		if isSet := cmd.Flags().Lookup("name").Changed; !isSet {
			cmd.Help()
			os.Exit(0)
		}
		groupStop(group)
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
