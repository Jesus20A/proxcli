package config

import (
	"fmt"
	"os"
	"proxcli/structs"

	"gopkg.in/yaml.v3"
)

var home, _ = os.UserHomeDir()

var Inventoryfile string = fmt.Sprintf("%s/.proxcli/inventory.yml", home)

// Obtain config parameters from config file
func InitConfig() (config map[string]string) {
	var configfile string = fmt.Sprintf("%s/.proxcli/proxcli.yml", home)

	data, err := os.ReadFile(configfile)
	if err != nil {
		fmt.Println(err)
	}

	c := structs.Config{}
	if err := yaml.Unmarshal(data, &c); err != nil {
		fmt.Println(err)
	}

	config = make(map[string]string)
	for _, value := range c.Config {
		config["node"] = value.Node
		config["ip"] = value.Ip
		config["user"] = value.Security.User
		config["realm"] = value.Security.Realm
		config["tokenid"] = value.Security.Tokenid
		config["token"] = value.Security.Token

	}

	return config

}
