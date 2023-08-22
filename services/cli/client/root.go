/*
Copyright © 2023 todo-list <dorlib318@gmail.com>
*/

package client

import (
	"github.com/spf13/cobra"
	"os"
)

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   "todo",
	Short: "TODO is a simple todo list app to manage your tasks",
	Long: `TODO is a CLI library for Go that lets you organize you tasks.
This application is a tool to track and update daily tasks and group's tasks.
With todo-list-manager you can make your day and even your group more organized and efficient.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(client *cobra.Command, args []string) { },
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Version: "1.0.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().String("key", "", "api-key for authentication")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
