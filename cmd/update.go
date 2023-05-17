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

// updateCmd represents the update command.
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update command will add update the info of a given task.",
	Long: ` update command will add update the info of a given task.
			update must except on of the following tags and inputs: 
			-t: the title of the task (accept string).
			-i: the ID of the task (accept string).
			
			in addition, it can accept one or more of the following tags:
			-t: the title of the task (accept string).
			-d: a shot description of the task (accept string).
			-p: the priority of the task, which can be: Critical, VeryHigh, High, Medium, Low (accept string).
			-l: the deadline of the task, in the following format: "dd/mm/yyyy (accept string)."
			-n: update the user who's the task is assigned to by name.
			-i: update the the user who's the task is assigned to by id.

`,
	Example: `todo update -t="homework" -d="do homework 3 in intro to cs"`,
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
	Run: func(cmd *cobra.Command, args []string) {
		flagsMap := make(map[string]interface{}, len(flagsUsed))

		for _, flag := range flagsUsed {
			flagsMap[flag] = cmd.Flag(flag).Value
		}

		// username := cmd.Flag("userid").Value.String()

		fmt.Println(flagsMap)

		if !Contains(flagsUsed, "task-id") && !Contains(flagsUsed, "task-title") {
			fmt.Println("you must specify task id or task name")

		}

		if Contains(flagsUsed, "task-id") {
			data.UpdateTaskByID((flagsMap["taskid"]).(uint), flagsMap)

			return
		}

		// try to find the task by the given title
		_, rowsAffected := data.GetTasksByTitle(flagsMap["task-title"].(string))
		if rowsAffected > 1 {
			fmt.Printf("more than one tasks exists with the title: %v, please specify task id", flagsMap["task-title"].(string))

			return
		}

		data.UpdateTaskByTitle(flagsMap["task-title"].(string), flagsMap)
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	RootCmd.AddCommand(updateCmd)
	updateCmd.PersistentFlags().Uint("task-id", 0, "the id of the task to update")
	updateCmd.PersistentFlags().String("task-title", "", "the id of the task to update")

	updateCmd.PersistentFlags().StringP("title", "t", "", "update the task's title")
	updateCmd.PersistentFlags().StringP("description", "d", "", "update the task's description")
	updateCmd.PersistentFlags().StringP("priority", "p", "", "update the task's priority")
	updateCmd.PersistentFlags().StringP("deadline", "l", "", "update the task's deadline")
	updateCmd.PersistentFlags().StringP("username", "n", "", "update the user who's the task is assigned to by name")
	updateCmd.PersistentFlags().UintP("userid", "i", 0, "update the the user who's the task is assigned to by id")

	updateCmd.Flags().SetInterspersed(false)
	updateCmd.MarkFlagsMutuallyExclusive("username", "userid")
}
