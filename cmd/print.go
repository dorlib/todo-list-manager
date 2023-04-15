/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"strconv"
	"todo/data"

	"github.com/spf13/cobra"
)

var taskTitle string

// printCmd represents the print command.
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "print command will print the given task/s to the console.",
	Long: ` print command will print the given task/s to the console.
			print must except one and only one from the following tags: 
			-a: print all the tasks to the console.
			-d: print all the tasks to the console, ordered by deadline (closest deadline first)
			-p: print all the tasks to the console, ordered by priority (most urgent first)
			-c: print all the tasks to the console, ordered by creation date (oldest creation date first)
			-i: print the task with the given ID to the console.
			-t: print the task with the given title to the console (if there is only one task with the given title).
			
			for example: 
			todo print -i="134", 
			todo print -a
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stringTaskID, err := cmd.Flags().GetString("ID")
		if err != nil {
			fmt.Printf("error while parsing flag: %v", err)
		}

		taskID, err := strconv.ParseUint(stringTaskID, 10, 64)
		if err != nil {
			fmt.Printf("err while parsing ID: %v", err)
			return
		}

		if rootCmd.Flags().Lookup("all") != nil {
			data.PrintAllTasks()
		} else if rootCmd.Flags().Lookup("by-deadline") != nil {

		} else if rootCmd.Flags().Lookup("by-priority") != nil {

		} else if rootCmd.Flags().Lookup("by-created-at") != nil {

		} else if taskTitle == "" {
			if !data.TaskExistByName(taskTitle) {
				fmt.Printf("Task %v Does Not Exist", taskTitle)
			} else {
				data.PrintTaskByName(taskTitle)
			}
		} else if stringTaskID != "" {
			if !data.TaskExistByID(uint(taskID)) {
				fmt.Printf("Task %v Does Not Exist", taskID)
			} else {
				data.PrintTaskByID(uint(taskID))
			}
		} else {
			if data.TaskExistByName(taskTitle) || data.TaskExistByID(uint(taskID)) {
				fmt.Println("Task Does Not Exist ")
			}

			data.PrintTaskByID(uint(taskID))
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
	printCmd.PersistentFlags().StringP("all", "a", "", "print all tasks")
	printCmd.PersistentFlags().StringP("by-deadline", "d", "", "print all tasks by order of deadline")
	printCmd.PersistentFlags().StringP("by-priority", "p", "", "print all tasks by priority")
	printCmd.PersistentFlags().StringP("by-created-at", "c", "", "print all tasks by order of time of creation")
	printCmd.PersistentFlags().StringP("ID", "i", "", "print task by ID")
	printCmd.PersistentFlags().StringVarP(&taskTitle, "title", "t", "", "print task by name")
}
