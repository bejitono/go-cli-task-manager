package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		fmt.Printf("added \"%s\" to you task list \n", task)
	},
}

// All init functions get run before the main function
func init() {
	RootCmd.AddCommand(addCmd)
}
