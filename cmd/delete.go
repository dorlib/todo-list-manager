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

// deleteCmd represents the delete command.
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete command will remove the given task from your todo list.",
	Long: `delete command will remove the given task from your todo list.
			delete must except one and only one from the following tags: 
			-t: the task's title (accept string).
			-i: the task's ID (accept string).
			-a: remove all the tasks in the list.
			
			for example: 
			todo delete -i="134", 
			todo delete -a
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")

		title, err := cmd.Flags().GetString("taskTitle")
		if err != nil {
			fmt.Printf("error while parsing flag: %v", err)
		}

		taskID, err := cmd.Flags().GetString("taskID")
		if err != nil {
			fmt.Printf("error while parsing flag: %v", err)
		}

		if title != "" {
			if !data.TaskExistByName(title) {
				fmt.Printf("Task %v Does Not Exist", taskTitle)

				return
			}

			data.DeleteTaskByTitle(title)
		} else if taskID != "" {
			if !data.TaskExistByName(taskID) {
				fmt.Printf("Task %v Does Not Exist", taskID)

				return
			}

			data.DeleteTaskByID(uuid.MustParse(taskID))
		}

	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	printCmd.LocalNonPersistentFlags().StringP("title", "t", "", "delete by the task's title")
	printCmd.LocalNonPersistentFlags().StringP("taskID", "i", "", "delete by the task's ID")
	printCmd.LocalNonPersistentFlags().StringP("all", "a", "", "delete all tasks")

	rootCmd.AddCommand(deleteCmd)
}
