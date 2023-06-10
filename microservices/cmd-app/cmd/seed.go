/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"todo/data"
	"todo/middlewares"
)

// seedCmd represents the seed command.
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed command will seed pre made data",
	Long: `Seed command will seed pre made data
			Seed will create users, groups and tasks which will be assigned to the different users.`,
	PersistentPreRunE: middlewares.AuthenticationMiddleware,
	Run: func(cmd *cobra.Command, args []string) {
		err := data.Seeder()
		if err != nil {
			fmt.Printf("error while seeding data: %v", err)

			return
		}
	},
}

func init() {
	RootCmd.AddCommand(seedCmd)
}
