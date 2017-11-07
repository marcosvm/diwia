package cmd

import (
	"fmt"
	"os"

	"github.com/slomek/diwia/clothes"
	"github.com/slomek/diwia/database"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds piece of clothes",
	Long:  "Adds a piece of clothing to the database",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		desc := args[0]
		if len(args) > 1 {
			desc = args[1]
		}

		db := database.OpenWrite()
		defer db.Close()

		if err := clothes.AddItem(db, name, desc); err != nil {
			fmt.Printf("Failed to insert new item: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("A new item has been inserted")
	},
}
