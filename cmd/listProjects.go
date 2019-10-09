/*
Copyright © 2019 Kondukto

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listProjectsCmd represents the listProjects command
var listProjectsCmd = &cobra.Command{
	Use:   "listProjects",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listProjects called")
	},
}

func init() {
	listCmd.AddCommand(listProjectsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listProjectsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listProjectsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
