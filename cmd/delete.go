/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"github.com/google/uuid"
	"todo/data"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command.
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")

		taskTitle, err := cmd.Flags().GetString("taskTitle")
		if err != nil {
			fmt.Printf("error while parsing flag: %v", err)
		}

		taskID, err := cmd.Flags().GetString("taskID")
		if err != nil {
			fmt.Printf("error while parsing flag: %v", err)
		}

		if taskTitle != "" {
			data.DeleteTaskByTitle(taskTitle)
		}

		if taskID != "" {
			data.DeleteTaskByID(uuid.MustParse(taskID))
		}

	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	printCmd.PersistentFlags().StringP("taskTitle", "t", "", "a"+"delete by the task's title")
	printCmd.PersistentFlags().StringP("taskID", "i", "", "delete by the task's ID")
	printCmd.PersistentFlags().StringP("all", "a", "", "delete all tasks")

	rootCmd.AddCommand(deleteCmd)
}
