package cmd

import (
	"fmt"
	"os"

	"github.com/slomek/diwia/clothes"
	"github.com/slomek/diwia/database"
	"github.com/spf13/cobra"
)

var wearCmd = &cobra.Command{
	Use:   "wear",
	Short: "Adds to clothes' wear number",
	Long:  "Increments a wear number for the piece of clothing",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		db := database.OpenWrite()
		defer db.Close()

		if err := clothes.IncWears(db, name); err != nil {
			fmt.Printf("Failed to update item: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("A wear-number for item has been updated")
	},
}
