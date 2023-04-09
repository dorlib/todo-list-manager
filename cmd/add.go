/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"github.com/spf13/cobra"
	"strings"
	"todo/data"
)

var title string
var description string
var priority string
var deadline string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		deadlineDate := strings.Split(deadline, "/")
		data.CreateTask(title, description, priority, data.Date{Day: deadlineDate[0], Month: deadlineDate[1], Year: deadlineDate[2]})
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	printCmd.PersistentFlags().StringVarP(&taskTitle, "title", "t", "", "add the task's title")
	printCmd.PersistentFlags().StringVarP(&taskTitle, "description", "d", "", "add the task's description")
	printCmd.PersistentFlags().StringVarP(&taskTitle, "priority", "p", "", "add the task's priority")
	printCmd.PersistentFlags().StringVarP(&taskTitle, "deadline", "d", "", "add the task's deadline")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
