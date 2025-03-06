/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a task by its ID.",
	Long:  `This function removes a task from the todo list based on its index.  If the ID is valid, the task is removed by slicing the Todos list and excluding the selected task. The updated list is then saved to the storage. If the ID is out of range, it returns an error." `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("No id is provided")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("id is not numberic")
			return
		}

		delete(id)
	},
}

func init() {

	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
