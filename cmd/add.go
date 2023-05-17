/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"strings"
	"todo/data"
)

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
		flags := ""
		cmd.Flags().Visit(func(flag *pflag.Flag) {
			flags += flag.Name + ","
		})

		flagsUsed = strings.Split(flags, ",")
		flagsUsed = flagsUsed[:len(flagsUsed)-1]

		fmt.Println(flagsUsed)

		return nil
	},

	Example: `todo add -t="homework" -d="do homework 3 in intro to cs" -p="High" -l="04/05/2023"`,
	Run: func(cmd *cobra.Command, args []string) {
		flagsMap := make(map[string]string, len(flagsUsed))

		for _, flag := range flagsUsed {
			flagsMap[flag] = cmd.Flag(flag).Value.String()
		}

		// username := cmd.Flag("userid").Value.String()

		fmt.Println(flagsMap)

		deadlineDate := strings.Split(flagsMap["deadline"], "/")
		data.CreateTask(flagsMap["title"], flagsMap["description"], flagsMap["priority"], data.Date{DeadlineDay: deadlineDate[0], DeadlineMonth: deadlineDate[1], DeadlineYear: deadlineDate[2]})
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	updateCmd.Flags().SetInterspersed(false)

	addCmd.PersistentFlags().StringP("title", "t", "", "add the task's title")
	addCmd.PersistentFlags().StringP("description", "d", "", "add the task's description")
	addCmd.PersistentFlags().StringP("priority", "p", "", "add the task's priority")
	addCmd.PersistentFlags().StringP("deadline", "l", "", "add the task's deadline")
	addCmd.MarkFlagsRequiredTogether("title", "description", "priority", "deadline")
}
