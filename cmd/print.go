/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"todo/data"

	"github.com/spf13/cobra"
)

var username string
var userid uint

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

			in addition, you can print a specific task by:
			-i: print the task with the given ID to the console (available only if the -u tag has been used).
			-t: print the task with the given title to the console (available only if the -u tag has been used).
`,
	Example: `todo print -i="134"
			  todo print --username=dor -a
			  todo print --userID=12 -t="fix bugs"	
`,

	Run: func(cmd *cobra.Command, args []string) {
		user, found := data.GetUser(userid, username)
		fmt.Println(user)

		if found {
			fmt.Println("replace..")
		} else {
			fmt.Println("replace..")
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
	printCmd.PersistentFlags().UintP("ID", "i", 0, "print task by ID")
	printCmd.PersistentFlags().StringP("title", "t", "", "print task by name")

	rootCmd.Flags().Lookup("by-deadline").NoOptDefVal = "deadline"
	rootCmd.Flags().Lookup("by-priority").NoOptDefVal = "priority"
	rootCmd.Flags().Lookup("by-created-at").NoOptDefVal = "created-at"
	rootCmd.Flags().Lookup("all").NoOptDefVal = "all"
}
