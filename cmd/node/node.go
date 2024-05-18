package node

import (
	"encoding/json"
	"fmt"
	"log"
	"proxcli/cmd"
	"proxcli/pkg/colors"
	"proxcli/pkg/config"
	"proxcli/pkg/filter"
	"proxcli/pkg/request"
	"proxcli/pkg/types"
	"strconv"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
)

func uptimeFormatter(uptime int) string {
	days := uptime / 86400
	hours := (uptime % 86400) / 3600
	minutes := (uptime % 3600) / 60
	return fmt.Sprintf("%d days %d hrs %d min", days, hours, minutes)
}

func nodeinfo() {
	config := config.InitConfig()
	url := fmt.Sprintf("https://%s:8006/api2/json/nodes/%s/status", config["ip"], config["node"])
	data, _ := request.NewRequest(url, "GET")
	var info types.NodeInfo
	err := json.Unmarshal(data, &info)
	if err != nil {
		log.Fatal(err)
	}

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Name"},
			{Align: simpletable.AlignCenter, Text: "Cpus"},
			{Align: simpletable.AlignCenter, Text: "Mem-Total"},
			{Align: simpletable.AlignCenter, Text: "Mem-Used"},
			{Align: simpletable.AlignCenter, Text: "Mem-Free"},
			{Align: simpletable.AlignCenter, Text: "Avg-1min"},
			{Align: simpletable.AlignCenter, Text: "Avg-5min"},
			{Align: simpletable.AlignCenter, Text: "Avg-15min"},
			{Align: simpletable.AlignCenter, Text: "Uptime"},
		},
	}

	var cells [][]*simpletable.Cell

	cells = append(cells, []*simpletable.Cell{
		{Text: colors.White(config["node"])},
		{Text: colors.White(strconv.Itoa(info.Node.Cpu.Cpus))},
		{Text: colors.White(filter.MemConverter(float32(info.Node.Memory["total"])))},
		{Text: colors.Yellow(filter.MemConverter(float32(info.Node.Memory["used"])))},
		{Text: colors.Green(filter.MemConverter(float32(info.Node.Memory["free"])))},
		{Text: colors.White(info.Node.Loadavg[0])},
		{Text: colors.Yellow(info.Node.Loadavg[1])},
		{Text: colors.Red(info.Node.Loadavg[2])},
		{Text: colors.Blue(uptimeFormatter(info.Node.Uptime))},
	})

	table.Body = &simpletable.Body{Cells: cells}
	table.SetStyle(simpletable.StyleRounded)
	table.Println()
}

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Interact with a node",
	Long:  `Perform actions on a node`,
}

var nodeGet = &cobra.Command{
	Use:   "get",
	Short: "Get node info",
	Long:  `Display more detail info about the node, like cpu and memory and load average`,
	Run: func(cmd *cobra.Command, args []string) {
		nodeinfo()
	},
}

func init() {
	cmd.RootCmd.AddCommand(nodeCmd)
	nodeCmd.AddCommand(nodeGet)
}
