/*
Copyright © 2024 Jesus Salas
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "proxcli",
	Short: "CLI clien for Proxmox API",
	Long: `A small CLI client for the Proxmox API that allow you to start and 
stop VMs, list the VMs and filter by status and get more detail info of VMs`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
