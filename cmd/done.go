/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"todo/data"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command.
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "done command marks the given task as done.",
	Long: `done command will marks the given task as done.
			done must except one and only one from the following tags: 
			-t: the task's title (accept string).
			-i: the task's ID (accept string).
`,
	Example: `todo done -i "134"`,
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
			data.ToggleDoneByTitle(taskTitle, true)
		} else if taskID != 0 {
			data.ToggleDoneByID(taskID, true)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	doneCmd.PersistentFlags().UintP("ID", "i", 0, "mark task by ID as done")
	doneCmd.PersistentFlags().StringP("title", "t", "", "mark task by title as done")
}
