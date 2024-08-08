/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package network

import (
	"infraTools/cmd/network/geolocate"
	"infraTools/cmd/network/nmap"
	"github.com/spf13/cobra"
)

// NetworkCmd represents the network command
var NetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "Network is a palette that contains based system commands",
	Long: `
	Network

The Network Tools module is designed to provide a comprehensive suite of utilities for network analysis and diagnostics.
This module includes commands that enable users to perform various network-related tasks efficiently.
Whether you need to scan networks, check connectivity, or gather information about network devices, 
this module has you covered.`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func addSubcommandPalettes() {
	NetworkCmd.AddCommand(nmap.PublicCmd)
	NetworkCmd.AddCommand(nmap.PrivateCmd)
	NetworkCmd.AddCommand(geolocate.GeolocateCmd)
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// networkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// networkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubcommandPalettes()
}
