/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"github.com/google/uuid"
	"todo/data"

	"github.com/spf13/cobra"
)

var TaskTitle string

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var taskName string
		var taskID uuid.UUID

		if rootCmd.Flags().Lookup("all") != nil {

		} else if rootCmd.Flags().Lookup("by-deadline") != nil {

		} else if rootCmd.Flags().Lookup("by-priority") != nil {

		} else if rootCmd.Flags().Lookup("by-created-at") != nil {

		} else if TaskTitle == "" {
			taskID := uuid.MustParse(args[0])
			if data.TaskExistByID(taskID) {
				fmt.Println("Task Does Not Exist ")
			}

			data.PrintTask(taskID)
		} else {
			if data.TaskExistByName(taskName) {
				fmt.Println("Task Does Not Exist ")
			}

			data.PrintTask(taskID)
		}
	},
}

func init() {
	rootCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	printCmd.PersistentFlags().StringP("all", "a", "", "print all tasks")
	printCmd.PersistentFlags().StringP("by-deadline", "d", "", "print all tasks by order of deadline")
	printCmd.PersistentFlags().StringP("by-priority", "p", "", "print all tasks by priority")
	printCmd.PersistentFlags().StringP("by-created-at", "c", "", "print all tasks by order of time of creation")
	printCmd.PersistentFlags().StringVarP(&TaskTitle, "title", "t", "", "print all tasks")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
