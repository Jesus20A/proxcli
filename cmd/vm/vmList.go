package vm

import (
	"encoding/json"
	"fmt"
	"log"
	"proxcli/pkg/colors"
	"proxcli/pkg/config"
	"proxcli/pkg/request"
	"proxcli/pkg/types"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
)

// Function to obtain all the Vms in the node
func Vmsinfo(state string, mode string) (vms types.VmInventory) {
	config := config.InitConfig()
	url := fmt.Sprintf("https://%s:8006/api2/json/nodes/%s/qemu", config["ip"], config["node"])
	data, _ := request.NewRequest(url, "GET")
	var info types.VmsInfo
	err := json.Unmarshal(data, &info)
	if err != nil {
		log.Fatal(err)
	}
	switch {
	case state == "running" || state == "stopped":
		makeTable(info, state)
	case state == "false" && mode == "silent":
		item := types.VmInventory{}
		err := json.Unmarshal(data, &item)
		if err != nil {
			log.Fatal(err)
		}
		vms = item
	default:
		makeTable(info, state)
	}
	return vms
}

func makeTable(info types.VmsInfo, state string) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "id"},
			{Align: simpletable.AlignCenter, Text: "Name"},
			{Align: simpletable.AlignCenter, Text: "Status"},
		},
	}
	var cells [][]*simpletable.Cell

	for _, item := range info.Data {

		if item.Status == state {
			cells = append(cells, []*simpletable.Cell{
				{Text: fmt.Sprintf("%d", item.Vmid)},
				{Text: colors.Blue(item.Name)},
				{Text: colors.Color(state, item.Status)},
			})
		} else if state == "all" {
			cells = append(cells, []*simpletable.Cell{
				{Text: fmt.Sprintf("%d", item.Vmid)},
				{Text: colors.Blue(item.Name)},
				{Text: colors.Color(item.Status, item.Status)},
			})
		}
	}

	table.Body = &simpletable.Body{Cells: cells}
	table.SetStyle(simpletable.StyleRounded)
	table.Println()
}

var VmsList = &cobra.Command{
	Use:   "list",
	Short: "List all the VMs",
	Long:  `Display all VMs along with their current status`,
	Run: func(cmd *cobra.Command, args []string) {
		state, _ := cmd.Flags().GetString("state")
		Vmsinfo(state, "verbose")
	},
}
