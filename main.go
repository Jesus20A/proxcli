package main

import (
	"proxcli/cmd"
	_ "proxcli/cmd/configure"
	_ "proxcli/cmd/inventory"
	_ "proxcli/cmd/lxc"
	_ "proxcli/cmd/node"
	_ "proxcli/cmd/vm"
)

func main() {
	cmd.Execute()
}
