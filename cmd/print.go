/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"github.com/spf13/pflag"
	"strconv"
	"todo/data"

	"github.com/spf13/cobra"
)

// todo - verify that user inserted flags according the rules described in the long description.

var username string
var userid uint

// printCmd represents the print command.
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "print command will print the given task/s to the console.",
	Long: ` print command will print the given task/s to the console.
			if no tags were added, print will print all the tasks.
			print can accept one and only one from the following optional tags:
			
			you can make the print command to look only on the tasks of a specific user, with:
			--username: the user's name.
			or
			--userid: the user's id.
			if none of them was not mentioned, the print command will look on all the tasks of the group.

			in addition, print can except one and only one from the following tags:
			-a: print all the tasks to the console.
			-d: print all the tasks to the console, ordered by deadline (closest deadline first)
			-p: print all the tasks to the console, ordered by priority (most urgent first)
			-c: print all the tasks to the console, ordered by creation date (oldest creation date first)

			you can also add the following tags even after selecting one of the optional tags above:
			--done: print only the done tasks.
			--undone: print only the undone tasks.
			--with-priority: print only the tasks with a specific priority.

			in addition, you can print a specific task by (cannot be with any of the tags above):
			-i: print the task with the given ID to the console (available only if the -u tag has been used).
			-t: print the task with the given title to the console (available only if the -u tag has been used).
	`,
	Example: `todo print -i="134"
			  todo print --username=dor -a
			  todo print --userID=12 -t="fix bugs"	
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			afterDash := pflag.Args()[pflag.CommandLine.ArgsLenAtDash():]
			fmt.Printf("args after dash: %v\n", afterDash)
		}
		fmt.Println(len(args))
		fmt.Println(args)

		taskID, err := cmd.Flags().GetUint("id")
		if err != nil {
			fmt.Printf("error while parsing flag: %v", err)
		}

		taskTitle, err := cmd.Flags().GetString("title")
		if err != nil {
			fmt.Printf("error while parsing flag: %v", err)
		}

		if len(args) == 0 {
			data.PrintAllTasks(data.User{}, false)

			return
		}

		if len(args) > 3 {
			fmt.Printf("accepts at most 3 arg(s), found: %v", len(args))

			return
		}

		if taskTitle != "" {
			if len(args) > 1 {
				fmt.Printf("accepts at most 3 arg(s), found: %v", len(args))

				return
			}

			data.PrintTaskByName(taskTitle)

			return
		} else if taskID != 0 {
			if len(args) > 1 {
				fmt.Printf("accepts at most 3 arg(s), found: %v", len(args))

				return
			}

			data.PrintTaskByID(taskID)

			return
		}

		user, found := data.GetUser(userid, username)

		byDeadLine, err := cmd.Flags().GetString("by-deadline")
		if err != nil {
			fmt.Printf("error while parsing flag: %v", err)
		}

		byPriority, err := cmd.Flags().GetString("by-priority")
		if err != nil {
			fmt.Printf("error while parsing flag: %v", err)
		}

		byCreatedAt, err := cmd.Flags().GetString("by-created-at")
		if err != nil {
			fmt.Printf("error while parsing flag: %v", err)
		}

		all, err := cmd.Flags().GetString("all")
		if err != nil {
			fmt.Printf("error while parsing flag: %v", err)
		}

		done, err := cmd.Flags().GetString("done")
		if err != nil {
			fmt.Printf("error while parsing flag: %v", err)
		}

		undone, err := cmd.Flags().GetString("undone")
		if err != nil {
			fmt.Printf("error while parsing flag: %v", err)
		}

		check := validatePrintTags(args)
		fmt.Println(check)

		if username != "" {
		}
		if byDeadLine == "deadline" {
			data.PrintByDeadLine(user, found)
		} else if byPriority == "priority" {
			data.PrintByPriority(user, found)
		} else if byCreatedAt == "created-at" {
			data.PrintByCreationDate(user, found)
		} else if all == "all" {
			data.PrintAllTasks(user, found)
		} else if done == "done" {
			fmt.Println(done)
		} else if undone == "undone" {
			fmt.Println(undone)
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
	printCmd.PersistentFlags().StringVar(&username, "username", "", "look on the tasks of a specific user name")
	printCmd.PersistentFlags().UintVar(&userid, "userid", 0, "look on the tasks of a specific user ID")

	printCmd.PersistentFlags().StringP("all", "a", "", "print all the tasks")
	printCmd.PersistentFlags().StringP("by-deadline", "d", "", "print all tasks by order of deadline")
	printCmd.PersistentFlags().StringP("by-priority", "p", "", "print all tasks by priority")
	printCmd.PersistentFlags().StringP("by-created-at", "c", "", "print all tasks by order of time of creation")

	printCmd.PersistentFlags().String("done", "", "print all the done tasks")
	printCmd.PersistentFlags().String("undone", "", "print all the undone tasks")
	printCmd.PersistentFlags().String("with-priority", "", "print all the done tasks with a given priority")

	printCmd.PersistentFlags().UintP("id", "i", 0, "print task by ID")
	printCmd.PersistentFlags().StringP("title", "t", "", "print task by name")

	//printCmd.Flags().Lookup("by-deadline").NoOptDefVal = "deadline"
	//printCmd.Flags().Lookup("by-priority").NoOptDefVal = "priority"
	//printCmd.Flags().Lookup("by-created-at").NoOptDefVal = "created-at"
	//printCmd.Flags().Lookup("all").NoOptDefVal = "all"
	//printCmd.Flags().Lookup("done").NoOptDefVal = "done"
	//printCmd.Flags().Lookup("undone").NoOptDefVal = "undone"
}

func validatePrintTags(args []string) bool {
	if len(args) > 3 {
		return false
	}

	if username != "" && userid != 0 {
		return false
	}

	if contains(args, username) || contains(args, strconv.Itoa(int(userid))) {

	}

	return true
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
