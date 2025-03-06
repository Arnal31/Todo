/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	id    int
	state bool
)

// toggleCmd represents the toggle command
var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Marks a task as completed or not completed.",
	Long:  `This function toggles the completion status of a task. If state is true, it marks the task as completed and sets the CompletedAt timestamp to the current time. If state is false, it marks the task as not completed and clears the CompletedAt field.`,
	Run: func(cmd *cobra.Command, args []string) {
		toggle(id, state)
	},
}

func init() {

	toggleCmd.Flags().IntVarP(&id, "id", "i", 0, "id of the task")
	toggleCmd.Flags().BoolVarP(&state, "state", "s", false, "state of the task")
	toggleCmd.MarkFlagRequired("id")
	toggleCmd.MarkFlagRequired("state")

	rootCmd.AddCommand(toggleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toggleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toggleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
