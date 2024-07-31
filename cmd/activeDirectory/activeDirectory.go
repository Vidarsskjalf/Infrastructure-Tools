/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package activeDirectory

import (
	"github.com/spf13/cobra"
)

// ActiveDirectoryCmd represents the activeDirectory command
var ActiveDirectoryCmd = &cobra.Command{
	Use:   "activeDirectory",
	Short: "activeDirectory is a palette that contains Active Directory based commands",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// activeDirectoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// activeDirectoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
