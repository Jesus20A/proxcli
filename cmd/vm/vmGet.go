package vm

import (
	"encoding/json"
	"fmt"
	"os"
	"proxcli/colors"
	"proxcli/config"
	"proxcli/filter"
	"proxcli/request"
	"proxcli/structs"
	"strconv"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
)

func Table(data []structs.VmInfo) {
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

func vmGet(name string, id int) {
	switch {
	case name != "none":
		exist := filter.Exist(name, group, id)
		if !exist {
			fmt.Printf("\u274C ERROR: No Vm with name %s found\n", colors.Red(name))
		} else {
			id = filter.GetId(name)
			info := filter.Vminfo(id)
			vm := []structs.VmInfo{info}
			Table(vm)

		}
	case id != 0:
		config := config.InitConfig()
		url := fmt.Sprintf("https://%s:8006/api2/json/nodes/%s/qemu/%s/status/current", config["ip"], config["node"], strconv.Itoa(id))
		data, code := request.NewRequest(url, "GET")
		if code == 200 {
			var info structs.VmInfo
			err := json.Unmarshal(data, &info)
			if err != nil {
				fmt.Println(err)
			}
			vm := []structs.VmInfo{info}
			Table(vm)
		} else {
			fmt.Printf("\u274C ERROR: No Vm with id %d found\n", id)
		}
	}
}

var VmGet = &cobra.Command{
	Use:   "get",
	Short: "Get Vm info",
	Long:  `Display more detail info about the Vm, like cpu and memory`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		id, _ := cmd.Flags().GetInt("id")
		if name == "none" && id == -1 {
			cmd.Help()
			os.Exit(0)
		}
		vmGet(name, id)
	},
}
