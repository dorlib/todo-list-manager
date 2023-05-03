/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"todo/data"

	"github.com/spf13/cobra"
)

// undoneCmd represents the undone command.
var undoneCmd = &cobra.Command{
	Use:   "undone",
	Short: "undone command marks the given task as done.",
	Long: `undone command will marks the given task as done.
			undone must except one and only one from the following tags: 
			-t: the task's title (accept string).
			-i: the task's ID (accept string).
`,
	Example: `todo undone -i="134"`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := cmd.Flags().GetUint("ID")
		if err != nil {
			fmt.Printf("error while parsing flag: %v", err)
		}

		taskTitle, err := cmd.Flags().GetString("title")
		if err != nil {
			fmt.Printf("error while parsing flag: %v", err)
		}

		if taskTitle != "" {
			if !data.TaskExistByName(taskTitle) {
				fmt.Printf("Task %v Does Not Exist", taskTitle)

				return
			}

			data.ToggleDoneByTitle(taskTitle, false)
		} else if taskID != 0 {
			if !data.TaskExistByID(taskID) {
				fmt.Printf("Task %v Does Not Exist", taskID)

				return
			}

			data.ToggleDoneByID(taskID, false)
		}
	},
}

func init() {
	RootCmd.AddCommand(undoneCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	undoneCmd.PersistentFlags().UintP("ID", "i", 0, "mark task by ID as undone")
	undoneCmd.PersistentFlags().StringP("title", "t", "", "mark task by title as undone")
}
