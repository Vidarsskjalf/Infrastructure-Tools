/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package system

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

// AdminsCmd represents the domainAdmins command
var AdminsCmd = &cobra.Command{
	Use:   "admins",
	Short: "A brief description of your command",
	Long: `
	AD Domain admins
	`,
	Run: func(cmd *cobra.Command, args []string) {
		scriptPath := "scripts/activeDirectory/admins.sh"
		out, err := exec.Command("bin/bash", scriptPath).Output()
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
	// domainAdminsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// domainAdminsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
