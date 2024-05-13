package lxc

import (
	"log"
	"os"
	"proxcli/pkg/colors"
	"proxcli/pkg/filter"
	"proxcli/pkg/types"
	"strconv"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
)

func Table(data types.LxcInfo) {
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
	cells = append(cells, []*simpletable.Cell{
		{Text: colors.White(strconv.Itoa(int((data.Data.Vmid.(float64)))))},
		{Text: colors.Blue(data.Data.Name)},
		{Text: colors.White(strconv.Itoa(data.Data.Cpus))},
		{Text: colors.Yellow(filter.MemConverter(float32(data.Data.Memory)))},
		{Text: colors.Color(data.Data.Status, data.Data.Status)},
	})

	table.Body = &simpletable.Body{Cells: cells}
	table.SetStyle(simpletable.StyleRounded)
	table.Println()

}

func lxcGet(id int, name string) {
	switch {
	case name != "none":
		id, err := filter.GetId(name, "lxc")
		if err != nil {
			log.Fatal(err)
		}
		info, _ := filter.Lxcinfo(id)
		Table(info)

	case id != 0:
		info, err := filter.Lxcinfo(id)
		if err != nil {
			log.Fatal(err)
		}
		Table(info)

	}
}

var LxcGet = &cobra.Command{
	Use:   "get",
	Short: "Get Lxc container info",
	Long:  `Display more detail info about the container, like cpu and memory`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		if name == "none" && id == -1 {
			cmd.Help()
			os.Exit(0)
		}
		lxcGet(id, name)
	},
}
