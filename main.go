/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"proxcli/cmd"
	_ "proxcli/cmd/node"
	_ "proxcli/cmd/vm"
)

func main() {
	cmd.Execute()
}
