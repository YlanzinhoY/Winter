/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spring_initializr/ylanzinhoy/internals/controller"
	dependencyorganization "github.com/spring_initializr/ylanzinhoy/internals/dependencyOrganization"
)

// apiExecuteCmd represents the apiExecute command
var apiExecuteCmd = &cobra.Command{
	Use:   "apiExecute",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		controller.ApiExecuteController()
		dependencyorganization.DeveloperTools()
	},
}

func init() {
	rootCmd.AddCommand(apiExecuteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiExecuteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiExecuteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
