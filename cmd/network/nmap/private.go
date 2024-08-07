/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package network

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

// PrivateCmd represents the private command
var PrivateCmd = &cobra.Command{
	Use:   "private",
	Short: "list all available private IP addresses",
	Long: `
	Private IP Tool
	
list all available private IP addresses using nmap`,
	Run: func(cmd *cobra.Command, args []string) {
		scriptPath := "scripts/nmap/private.sh"
		out, err := exec.Command("/bin/bash", scriptPath).Output()
		if err != nil {
			fmt.Printf("Error executing script: %v\n", err)
			return
		}
		fmt.Printf("Script output: %s\n", string(out))
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// privateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// privateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
