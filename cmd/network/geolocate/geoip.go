package geolocate

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
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
		ip := args[0]
		api_key := "Enter api key"
		url := fmt.Sprintf("https://ipinfo.io/%s?token=%s", ip, api_key)

		client := resty.New()
		resp, err := client.R().
			SetHeader("Accept", "application/json").
			Get(url)
		if err != nil {
			log.Fatalf("Error making request: %v", err)
		}

		var geoIP GeoResponse
		if err := json.Unmarshal(resp.Body(), &geoIP); err != nil {
			log.Fatalf("Error parsing response: %v", err)
		}

		fmt.Printf("IP: %s\n", geoIP.IP)
		fmt.Printf("City: %s\n", geoIP.City)
		fmt.Printf("Region: %s\n", geoIP.Region)
		fmt.Printf("Country: %s\n", geoIP.Country)
		fmt.Printf("Location: %s\n", geoIP.Loc)
		fmt.Printf("Organization: %s\n", geoIP.Org)
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
