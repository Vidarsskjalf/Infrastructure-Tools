/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package system

import (
	"infraTools/cmd/system/activeDirectory"
	"github.com/spf13/cobra"
)

// SystemCmd represents the system command
var SystemCmd = &cobra.Command{
	Use:   "system",
	Short: "System is a palette that contains based system commands",
	Long: `
	System

The system command allows you to perform various system-related tasks
such as managing Active Directory, checking system status, and more.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}


func addSubcommandPalettes() {
	SystemCmd.AddCommand(system.AdminsCmd)
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// systemCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// systemCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addSubcommandPalettes()
}
