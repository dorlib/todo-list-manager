/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"github.com/google/uuid"
	"strconv"
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
			
			for example: 
			todo undone -i="134", 
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stringTaskID, err := cmd.Flags().GetString("ID")
		if err != nil {
			fmt.Printf("error while parsing flag: %v", err)
		}

		taskTitle, err := cmd.Flags().GetString("title")
		if err != nil {
			fmt.Printf("error while parsing flag: %v", err)
		}

		taskID, err := strconv.ParseUint(stringTaskID, 10, 64)
		if err != nil {
			fmt.Printf("err while parsing ID: %v", err)

			return
		}

		if taskTitle != "" {
			if !data.TaskExistByName(taskTitle) {
				fmt.Printf("Task %v Does Not Exist", taskTitle)

				return
			}

			data.ToggleDoneByTitle(taskTitle, false)
		} else if stringTaskID != "" {
			if !data.TaskExistByID(uint(taskID)) {
				fmt.Printf("Task %v Does Not Exist", uuid.MustParse(stringTaskID))

				return
			}

			data.ToggleDoneByID(uint(taskID), false)
		}
	},
}

func init() {
	rootCmd.AddCommand(undoneCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	undoneCmd.PersistentFlags().StringP("ID", "i", "", "mark task by ID as undone")
	undoneCmd.PersistentFlags().StringP("title", "t", "", "mark task by title as undone")
}
