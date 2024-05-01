package filter

import (
	"encoding/json"
	"fmt"
	"os"
	"proxcli/config"
	"proxcli/request"
	"proxcli/structs"
	"strconv"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func Vminfo(id int) (info structs.VmInfo) {
	config := config.InitConfig()
	url := fmt.Sprintf("https://%s:8006/api2/json/nodes/%s/qemu/%s/status/current", config["ip"], config["node"], strconv.Itoa(id))
	data, _ := request.NewRequest(url, "GET")
	err1 := json.Unmarshal(data, &info)
	if err1 != nil {
		fmt.Println(err1)
	}
	return info
}

func WritetoFile(data string, filepath string) error {

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		cobra.CheckErr(err)
	}

	defer file.Close()

	_, err2 := file.WriteString(data)

	if err2 != nil {
		cobra.CheckErr(err2)
	}
	return err2
}

func Getgroup(group string) (hosts []structs.Vm) {

	data, err := os.ReadFile(config.Inventoryfile)
	if err != nil {
		fmt.Println(err)
	}

	c := structs.Groups{}
	if err := yaml.Unmarshal(data, &c); err != nil {
		fmt.Println(err)
	}

	for _, v := range c.Groups {
		if v.Name == group {
			hosts = v.Vms
			break
		}
	}
	return hosts
}

func GetId(name string) (id int) {
	data, err := os.ReadFile(config.Inventoryfile)
	if err != nil {
		fmt.Println(err)
	}
	c := structs.Vms{}
	if err := yaml.Unmarshal(data, &c); err != nil {
		fmt.Println(err)
	}
	for _, v := range c.Vms {
		if v.Name == name {
			id = v.Id
			break
		}
	}
	return id
}

func Exist(name string, group string, id int) (exist bool) {
	data, err := os.ReadFile(config.Inventoryfile)
	if err != nil {
		fmt.Println(err)
	}

	g := structs.Groups{}
	if err := yaml.Unmarshal(data, &g); err != nil {
		fmt.Println(err)
	}
	for _, v := range g.Groups {
		if v.Name == group {
			exist = true
			break
		}
	}

	n := structs.Vms{}
	if err := yaml.Unmarshal(data, &n); err != nil {
		fmt.Println(err)
	}
	for _, v := range n.Vms {
		if v.Name == name {
			exist = true
			break
		}
	}

	i := Vminfo(id)
	if i.Data.Vmid == id {
		exist = true
	}
	return exist
}

func MemConverter(bytes float32) string {
	GB := float32(bytes / (1024 * 1024 * 1024))
	return fmt.Sprintf("%.2f GB", GB)
}
