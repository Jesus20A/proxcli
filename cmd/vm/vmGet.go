package vm

import (
	"fmt"
	"log"
	"os"
	"proxcli/pkg/colors"
	"proxcli/pkg/filter"
	"proxcli/pkg/types"
	"strconv"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
)

func vmtable(data types.VmInfo) {
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
		{Text: fmt.Sprintf("%d", data.Data.Vmid)},
		{Text: colors.Blue(data.Data.Name)},
		{Text: colors.White(strconv.Itoa(data.Data.Cpus))},
		{Text: colors.Yellow(filter.MemConverter(float32(data.Data.Memory)))},
		{Text: colors.Color(data.Data.Status, data.Data.Status)},
	})

	table.Body = &simpletable.Body{Cells: cells}
	table.SetStyle(simpletable.StyleRounded)
	table.Println()

}

func vmGet(name string, id int) {
	switch {
	case name != "none":
		id, err := filter.GetId(name, "qemu")
		if err != nil {
			log.Fatal(err)
		}
		info, _ := filter.Vminfo(id)
		vmtable(info)

	case id != 0:
		info, err := filter.Vminfo(id)
		if err != nil {
			log.Fatal(err)
		}
		vmtable(info)

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
