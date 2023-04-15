/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"errors"
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
			-l: the deadline of the task, in the following format: "dd/mm/yyyy (accept string)."
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if title == "" || description == "" || priority == "" || deadline == "" {
			return errors.New("accepts 1 arg(s)")
		}
		return nil
	},
	Example: `todo add -t="homework" -d="do homework 3 in intro to cs" -p="High" -d="04/05/2023"`,
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
	addCmd.PersistentFlags().StringVarP(&title, "title", "t", "", "add the task's title")
	addCmd.PersistentFlags().StringVarP(&description, "description", "d", "", "add the task's description")
	addCmd.PersistentFlags().StringVarP(&priority, "priority", "p", "", "add the task's priority")
	addCmd.PersistentFlags().StringVarP(&deadline, "deadline", "l", "", "add the task's deadline")
}
