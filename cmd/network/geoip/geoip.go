/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package geoip

import (
	"github.com/spf13/cobra"
)

// GeoipCmd represents the geoip command
var GeoipCmd = &cobra.Command{
	Use:   "geoip",
	Short: "A tool to determine the geographical location of an ip address",
	Long: `
	IP Geolocation Tool
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// geoipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// geoipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
