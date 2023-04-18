/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"todo/data"

	"github.com/spf13/cobra"
)

var taskTitle string

// printCmd represents the print command.
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "print command will print the given task/s to the console.",
	Long: ` print command will print the given task/s to the console.
			print can accept one and only one from the following optional tags: 
			
			you can make the print command to look only on the tasks of a specific user, with:
			--username: the user's name.
			or
			--userid: the user's id.
			if none of them was not mentioned, the print command will look on all the tasks of the group.

			in addition, print must except one and only one from the following tags:
			-a: print all the tasks to the console.
			-d: print all the tasks to the console, ordered by deadline (closest deadline first)
			-p: print all the tasks to the console, ordered by priority (most urgent first)
			-c: print all the tasks to the console, ordered by creation date (oldest creation date first)

			in addition, this will always work:
			-i: print the task with the given ID to the console (available only if the -u tag has been used).
			-t: print the task with the given title to the console (available only if the -u tag has been used).
			

`,
	Example: `todo print -i="134"
			  todo print --username=dor -a
`,

	Run: func(cmd *cobra.Command, args []string) {
		if rootCmd.Flags().Lookup("all") != nil {
			data.PrintAllTasks()
		}

		taskID, err := cmd.Flags().GetUint("ID")
		if err != nil {
			fmt.Printf("error while parsing flag: %v", err)
		} else if rootCmd.Flags().Lookup("by-deadline") != nil {

		} else if rootCmd.Flags().Lookup("by-priority") != nil {

		} else if rootCmd.Flags().Lookup("by-created-at") != nil {

		} else if taskTitle == "" {
			if !data.TaskExistByName(taskTitle) {
				fmt.Printf("Task %v Does Not Exist", taskTitle)
			} else {
				data.PrintTaskByName(taskTitle)
			}
		} else if taskID != 0 {
			if !data.TaskExistByID(taskID) {
				fmt.Printf("Task %v Does Not Exist", taskID)
			} else {
				data.PrintTaskByID(taskID)
			}
		} else {
			if data.TaskExistByName(taskTitle) || data.TaskExistByID(taskID) {
				fmt.Println("Task Does Not Exist ")
			}

			data.PrintTaskByID(taskID)
		}
	},
}

func init() {
	rootCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	printCmd.PersistentFlags().StringP("by", "b", "", "input can be: 1. all 2.user")
	printCmd.PersistentFlags().StringP("name", "n", "", "the name of the user (comes after choosing option 2 in 'by' flag.")
	printCmd.PersistentFlags().StringP("by-deadline", "d", "", "print all tasks by order of deadline")
	printCmd.PersistentFlags().StringP("by-priority", "p", "", "print all tasks by priority")
	printCmd.PersistentFlags().StringP("by-created-at", "c", "", "print all tasks by order of time of creation")
	printCmd.PersistentFlags().UintP("ID", "i", 0, "print task by ID")
	printCmd.PersistentFlags().StringVarP(&taskTitle, "title", "t", "", "print task by name")

	rootCmd.Flags().Lookup("by-deadline").NoOptDefVal = "user"
}
