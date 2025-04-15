package cmd

import (
	"github.com/emms-garcia/golang-playground/todo-cli/internal/db"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "A simple CLI to manage your todo list",
	Long:  "todo-cli is a simple command line interface to manage your todo list.",
	CompletionOptions: cobra.CompletionOptions{
		// to disable the default command "completion"
		DisableDefaultCmd: true,
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		db, err := db.InitDatabase()
		if err != nil {
			return err
		}
		_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS todos (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				title TEXT NOT NULL,
				done BOOLEAN NOT NULL DEFAULT 0
			);
		`)
		return err
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(doneCmd)
	rootCmd.AddCommand(deleteCmd)
}
