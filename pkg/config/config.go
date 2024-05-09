package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Configuration Structs
type Security struct {
	User    string
	Realm   string
	Tokenid string
	Token   string
}
type Server struct {
	Node     string
	Ip       string
	Security Security
}
type Config struct {
	Config []Server
}

var home, _ = os.UserHomeDir()

var Inventoryfile string = fmt.Sprintf("%s/.proxcli/inventory.yml", home)

// Obtain config parameters from config file
func InitConfig() (config map[string]string) {
	var configfile string = fmt.Sprintf("%s/.proxcli/proxcli.yml", home)

	data, err := os.ReadFile(configfile)
	if err != nil {
		fmt.Println(err)
	}

	c := Config{}
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
