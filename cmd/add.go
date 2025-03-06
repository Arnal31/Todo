/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task to the todo list.",
	Long:  `This function creates a new task with the provided title, sets its Completed status to false, and assigns the current time to CreatedAt. The CompletedAt field is set to nil. The new task is then appended to the existing list of todos, which is saved to the storage.`,
	Run: func(cmd *cobra.Command, args []string) {
		if 1 > len(args) {
			fmt.Println("No tasks are provided")
			return
		}
		title := args[0]
		add(title)

	},
}

func init() {

	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
