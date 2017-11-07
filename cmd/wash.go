package cmd

import (
	"fmt"
	"os"

	"github.com/slomek/diwia/clothes"
	"github.com/slomek/diwia/database"
	"github.com/spf13/cobra"
)

var washCmd = &cobra.Command{
	Use:   "wash",
	Short: "Clears clothes' wear number",
	Long:  "Resets wear numbers for pieces of clothing to zero",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db := database.OpenWrite()
		defer db.Close()

		if err := clothes.ClearWears(db, args); err != nil {
			fmt.Printf("Failed to clear wear counts: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("A wear counts for item has been reset for %v\n", args)
	},
}
