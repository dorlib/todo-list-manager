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

// addCmd represents the add command.
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add command will add new task to your todo list.",
	Long: `add command will add new task to your todo list.
			Add must except the following tags and inputs: 
			-t: the title of the task (accept string).
			-d: a shot description of the task (accept string).
			-p: the priority of the task, which can be: Critical, VeryHigh, High, Medium, Low (accept string).
			-d: the deadline of the task, in the following format: "dd/mm/yyyy (accept string)."
			
			for example: 
			todo add -t="homework" -d="do homework 3 in intro to cs" -p="High" -d="04/05/2023"
`,
	Args: cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		deadlineDate := strings.Split(deadline, "/")
		data.CreateTask(title, description, priority, data.Date{Day: deadlineDate[0], Month: deadlineDate[1], Year: deadlineDate[2]})
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	printCmd.PersistentFlags().StringVarP(&taskTitle, "title", "t", "", "add the task's title")
	printCmd.PersistentFlags().StringVarP(&taskTitle, "description", "d", "", "add the task's description")
	printCmd.PersistentFlags().StringVarP(&taskTitle, "priority", "p", "", "add the task's priority")
	printCmd.PersistentFlags().StringVarP(&taskTitle, "deadline", "d", "", "add the task's deadline")
}
