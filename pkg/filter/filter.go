package filter

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"proxcli/pkg/colors"
	"proxcli/pkg/config"
	"proxcli/pkg/request"
	"proxcli/pkg/types"
	"strconv"

	"gopkg.in/yaml.v3"
)

func Vminfo(id int) (info types.VmInfo, err error) {
	config := config.InitConfig()
	url := fmt.Sprintf("https://%s:8006/api2/json/nodes/%s/qemu/%s/status/current", config["ip"], config["node"], strconv.Itoa(id))
	data, code := request.NewRequest(url, "GET")
	if code != 200 {
		err = fmt.Errorf("\u274C ERROR: No Vm with id %s found", colors.Red(strconv.Itoa(id)))
	} else {
		err := json.Unmarshal(data, &info)
		if err != nil {
			log.Fatal(err)
		}
	}
	return info, err
}

func Lxcinfo(id int) (info types.LxcInfo, err error) {
	config := config.InitConfig()
	url := fmt.Sprintf("https://%s:8006/api2/json/nodes/%s/lxc/%s/status/current", config["ip"], config["node"], strconv.Itoa(id))
	data, code := request.NewRequest(url, "GET")
	if code != 200 {
		err = fmt.Errorf("\u274C ERROR: No container with id %s found", colors.Red(strconv.Itoa(id)))
	} else {
		err := json.Unmarshal(data, &info)
		if err != nil {
			log.Fatal(err)
		}
	}
	return info, err
}

func Getgroup(group string) (hosts []types.Vm, err error) {

	data, err := os.ReadFile(config.Inventoryfile)
	if err != nil {
		log.Fatal(err)
	}

	g := types.Groups{}
	if err := yaml.Unmarshal(data, &g); err != nil {
		log.Fatal(err)
	}
	exist := false
	for _, v := range g.Groups {
		e := &exist
		if v.Name == group {
			hosts = v.Vms
			*e = true
			break
		}
	}

	if !exist {
		err = fmt.Errorf("\u274C ERROR: No Group with name %s found", colors.Red(group))
	}
	return hosts, err
}

func GetId(name, t string) (id int, err error) {
	data, err := os.ReadFile(config.Inventoryfile)
	if err != nil {
		log.Fatal(err)
	}
	switch t {
	case "qemu":
		v := types.Vms{}
		if err := yaml.Unmarshal(data, &v); err != nil {
			log.Fatal(err)
		}
		exist := false
		for _, vm := range v.Vms {
			e := &exist
			if vm.Name == name {
				id = vm.Id
				*e = true
				break
			}
		}
		if !exist {
			err = fmt.Errorf("\u274C ERROR: No Vm with name %s found", colors.Red(name))
		}
	case "lxc":
		l := types.LxcInventory{}
		if err := yaml.Unmarshal(data, &l); err != nil {
			log.Fatal(err)
		}
		exist := false
		for _, lxc := range l.Lxc {
			e := &exist
			if lxc.Name == name {
				id = lxc.Id.(int)
				*e = true
				break
			}
		}
		if !exist {
			err = fmt.Errorf("\u274C ERROR: No Container with name %s found", colors.Red(name))
		}
	}
	return id, err
}

func WritetoFile(data string, filepath string) (err error) {

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(data)

	if err != nil {
		return err
	}
	return err
}

func MemConverter(bytes float32) string {
	GB := float32(bytes / (1024 * 1024 * 1024))
	return fmt.Sprintf("%.2f GB", GB)
}
