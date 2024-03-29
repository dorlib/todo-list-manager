/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package user

import (
	"fmt"
	"todo/client"

	"github.com/spf13/cobra"
)

// connectCmd represents the connect command.
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("password validation...")
		fmt.Println("authentication...")
		fmt.Println("connect called")
	},
}

func init() {
	client.RootCmd.AddCommand(connectCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	connectCmd.PersistentFlags().StringP("username", "u", "", "username for connection")
	connectCmd.PersistentFlags().StringP("password", "p", "", "password for connection")
	connectCmd.MarkFlagsRequiredTogether("username", "password")
}
