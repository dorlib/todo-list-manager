/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
			-d: the deadline of the task, in the following format: "dd/mm/yyyy (accept string)."
			
			for example: 
			todo update -t="homework" -d="do homework 3 in intro to cs"
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
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
	rootCmd.AddCommand(updateCmd)
}
