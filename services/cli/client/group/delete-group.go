/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package group

import (
	"fmt"
	"todo/client"
	"todo/middlewares"

	"github.com/spf13/cobra"
)

// deleteGroupCmd represents the deleteGroup command.
var deleteGroupCmd = &cobra.Command{
	Use:   "delete-group",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRunE: middlewares.AuthenticationMiddleware,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteGroup called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteGroupCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteGroupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	client.RootCmd.AddCommand(deleteGroupCmd)
}
