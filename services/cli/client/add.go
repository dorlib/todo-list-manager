/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package client

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"log"
	"strconv"
	"strings"
	data2 "todo/data"
	"todo/middlewares"
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

			-u: assign the task to a user. if not specified, the task will be assigned to the creator.
			
			if u was specified, you also need to give one of the following: 
			--id: the id of the user to assign the task to.
			--name: the full name of the user to assign the task to.
			
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

	Example:           `todo add -t="homework" -d="do homework 3 in intro to cs" -p="High" -l="04/05/2023"`,
	PersistentPreRunE: middlewares.AuthenticationMiddleware,
	Run: func(cmd *cobra.Command, args []string) {
		flagsMap := make(map[string]string, len(flagsUsed))

		if Contains(flagsUsed, "user") && !Contains(flagsUsed, "id") && !Contains(flagsUsed, "name") {
			fmt.Printf("must specify user ID or username when assigning task to user")

			return
		}

		for _, flag := range flagsUsed {
			flagsMap[flag] = cmd.Flag(flag).Value.String()
		}

		// username := client.Flag("userid").Value.String()

		fmt.Println(flagsMap)

		id, err := strconv.ParseUint(flagsMap["id"], 10, 64)
		if err != nil {
			fmt.Println(err)
		}

		userToAssign, exists := data2.GetUser(uint(id), flagsMap["name"])

		if !exists {
			fmt.Printf("get the user who is currently connected")
		}

		deadlineDate := strings.Split(flagsMap["deadline"], "/")
		task, err := data2.CreateTask(flagsMap["title"], flagsMap["description"], flagsMap["priority"], data2.Date{DeadlineDay: deadlineDate[0], DeadlineMonth: deadlineDate[1], DeadlineYear: deadlineDate[2]}, userToAssign)
		if err != nil {
			fmt.Printf("%v", err)

			return
		}

		log.Printf("task created: %v", task)
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

	addCmd.PersistentFlags().StringP("user", "u", "user", "assign the task to a user")
	addCmd.PersistentFlags().Uint("id", 0, "the id of the user which the task will be assigned to")
	addCmd.PersistentFlags().String("name", "", "the name of the user which the task will be assigned to")
	addCmd.MarkFlagsMutuallyExclusive("id", "name")
}
