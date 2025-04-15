package cmd

import (
	"fmt"

	"github.com/emms-garcia/golang-playground/todo-cli/internal/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks on the to-do list",
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := db.GetDatabase()
		if err != nil {
			return fmt.Errorf("error connecting to database: %v", err)
		}
		defer db.Close()

		rows, err := db.Query("SELECT id, title, done FROM todos")
		if err != nil {
			return fmt.Errorf("error listing tasks: %v", err)
		}

		fmt.Println("Tasks:")
		for rows.Next() {
			var id int
			var title string
			var done bool
			rows.Scan(&id, &title, &done)
			status := " "
			if done {
				status = "✓"
			}
			fmt.Printf("[%s] %d: %s\n", status, id, title)
		}
		return nil
	},
}
