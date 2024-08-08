/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package geolocate

import (
	"github.com/spf13/cobra"
)


type GeoResponse struct {
	IP		string `json:"ip"`
	City	string `json:"city"`
	Region	string `json:"region"`
	Country	string `json:"country"`
	Loc		string `json:"loc"`
	Org		string `json:"org"`
}



// GeolocateCmd represents the geoip command
var GeolocateCmd = &cobra.Command{
	Use:   "geolocate [ip]",
	Short: "Geolocate an IP address",
	Long: `
	Geolocation Tool
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
