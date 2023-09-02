package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fiber",
	Short: "A colorful CLI tool",
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List items",
		Run: func(cmd *cobra.Command, args []string) {
			// Implement the list functionality here
			fmt.Println("Listing items...")
		},
	}

	var listMigrationsCmd = &cobra.Command{
		Use:   "list-migrations",
		Short: "List database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			// Implement the logic to list database migrations here
			fmt.Println("Migration items...")
		},
	}

	// Add the "list" command to the root command
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(listMigrationsCmd)
}
