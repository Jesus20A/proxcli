package start

import (
	"fmt"
	"log"
	"proxcli/pkg/config"
	"proxcli/pkg/filter"
	"proxcli/pkg/request"
	"strconv"
)

func Start(id int, name, t string) {
	config := config.InitConfig()
	base_url := fmt.Sprintf("https://%s:8006/api2/json/nodes/%s", config["ip"], config["node"])
	switch {
	case t == "lxc":
		switch {
		case name != "none":
			id, err := filter.GetId(name, "lxc")
			if err != nil {
				log.Fatal(err)
			}
			lxc, err := filter.Lxcinfo(id)
			if err != nil {
				log.Fatal(err)
			}
			if lxc.Data.Status == "running" {
				fmt.Printf("\u2705 %s Container is already running\n", lxc.Data.Name)
			} else {
				url := fmt.Sprintf("%s/%s/%s/status/start", base_url, t, strconv.Itoa(id))
				data, code := request.NewRequest(url, "POST")
				if code == 200 {
					fmt.Printf("\u2705 %s Container started\n", lxc.Data.Name)
				} else {
					fmt.Printf("\u274C %d - ERROR %s\n", code, data)
				}
			}

		default:
			lxc, err := filter.Lxcinfo(id)
			if err != nil {
				log.Fatal(err)
			}
			if lxc.Data.Status == "running" {
				fmt.Printf("\u2705 %s Container is already running\n", lxc.Data.Name)
			} else {
				url := fmt.Sprintf("%s/%s/%s/status/start", base_url, t, strconv.Itoa(id))
				data, code := request.NewRequest(url, "POST")
				if code == 200 {
					fmt.Printf("\u2705 %s Container started\n", lxc.Data.Name)
				} else {
					fmt.Printf("\u274C %d - ERROR %s\n", code, data)
				}
			}

		}
	case t == "qemu":
		switch {
		case name != "none":
			id, err := filter.GetId(name, "qemu")
			if err != nil {
				log.Fatal(err)
			}
			vm, err := filter.Vminfo(id)
			if err != nil {
				log.Fatal(err)
			}
			if vm.Data.Status == "running" {
				fmt.Printf("\u2705 %s Vm is already running\n", vm.Data.Name)
			} else {
				url := fmt.Sprintf("%s/%s/%s/status/start", base_url, t, strconv.Itoa(id))
				data, code := request.NewRequest(url, "POST")
				if code == 200 {
					fmt.Printf("\u2705 %s Vm started\n", vm.Data.Name)
				} else {
					fmt.Printf("\u274C %d - ERROR %s\n", code, data)
				}
			}

		default:
			vm, err := filter.Vminfo(id)
			if err != nil {
				log.Fatal(err)
			}
			if vm.Data.Status == "running" {
				fmt.Printf("\u2705 %s Vm is already running\n", vm.Data.Name)
			} else {
				url := fmt.Sprintf("%s/%s/%s/status/start", base_url, t, strconv.Itoa(id))
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
