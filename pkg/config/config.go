package config

import (
	"fmt"
	"log"
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
var Configdir string = fmt.Sprintf("%s/.proxcli", home)
var Inventoryfile string = fmt.Sprintf("%s/inventory.yaml", Configdir)
var Configfile string = fmt.Sprintf("%s/proxcli.yaml", Configdir)

func CreateConfig() {
	_, err := os.Stat(Configdir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("[*] No config directory found in %s\n", Configdir)
			fmt.Print("[*] Do you want to create the directory? [y/n] ")
			var answer string
			fmt.Scanln(&answer)
			switch answer {
			case "y":
				os.MkdirAll(Configdir, 0755)
				fmt.Printf("\u2705Configuration directory created at: %s\n", Configdir)
			case "n":
				fmt.Printf("[*] No config directory found in %s\n", Configdir)
				os.Exit(1)
			default:
				log.Fatal("\u274C Error: invalid input")
			}
		}
	} else {
		fmt.Printf("\u2705 Configuration directory already created at %s\n", Configdir)
	}
	fmt.Printf("[*] Creating configuration file at %v\n", Configfile)
	conf := Config{}
	fmt.Print("[*] Enter your proxmox node name: ")
	var node string
	fmt.Scanln(&node)
	fmt.Print("[*] Enter your proxmox ip address: ")
	var ip string
	fmt.Scanln(&ip)
	fmt.Print("[*] Enter your proxmox username: ")
	var user string
	fmt.Scanln(&user)
	var realm string = "pve"
	fmt.Print("[*] Enter your proxmox token id: ")
	var tokenid string
	fmt.Scanln(&tokenid)
	fmt.Print("[*] Enter your proxmox token: ")
	var token string
	fmt.Scanln(&token)
	conf.Config = append(conf.Config, Server{
		Node:     node,
		Ip:       ip,
		Security: Security{User: user, Realm: realm, Tokenid: tokenid, Token: token},
	})
	data, err := yaml.Marshal(conf)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(Configfile, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\u2705 Configuration file created at %s\n", Configfile)
}

// Obtain config parameters from config file
func InitConfig() (config map[string]string) {

	data, err := os.ReadFile(Configfile)
	if err != nil {
		log.Fatal(err)
	}

	c := Config{}
	if err := yaml.Unmarshal(data, &c); err != nil {
		log.Fatal(err)
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
