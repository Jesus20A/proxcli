package lxc

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
func Lxcsinfo(state, mode string) (lxc types.LxcInventory) {
	config := config.InitConfig()
	url := fmt.Sprintf("https://%s:8006/api2/json/nodes/%s/lxc", config["ip"], config["node"])
	data, _ := request.NewRequest(url, "GET")
	var info types.LxcsInfo
	err := json.Unmarshal(data, &info)
	if err != nil {
		log.Fatal(err)
	}
	switch {
	case state == "false" && mode == "silent":
		item := types.LxcInventory{}
		err := json.Unmarshal(data, &item)
		if err != nil {
			log.Fatal(err)
		}
		lxc = item
	default:
		listTable(info, state)
	}
	return lxc

}

func listTable(info types.LxcsInfo, state string) {
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
				{Text: colors.White(item.Vmid.(string))},
				{Text: colors.Blue(item.Name)},
				{Text: colors.Color(state, item.Status)},
			})
		} else if state == "all" {
			cells = append(cells, []*simpletable.Cell{
				{Text: colors.White(item.Vmid.(string))},
				{Text: colors.Blue(item.Name)},
				{Text: colors.Color(item.Status, item.Status)},
			})
		}
	}

	table.Body = &simpletable.Body{Cells: cells}
	table.SetStyle(simpletable.StyleRounded)
	table.Println()
}

var LxcList = &cobra.Command{
	Use:   "list",
	Short: "List all the Lxc containers",
	Long:  `Display all containers along with their current status`,
	Run: func(cmd *cobra.Command, args []string) {
		state, _ := cmd.Flags().GetString("state")
		Lxcsinfo(state, "verbose")
	},
}
