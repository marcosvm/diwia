package cmd

import (
	"fmt"
	"os"

	"github.com/slomek/diwia/clothes"
	"github.com/slomek/diwia/database"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all clothes",
	Long:  "Lists all pieces of clothing defined in the database",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		db := database.OpenRead()
		defer db.Close()

		pieces, err := clothes.GetAll(db)
		if err != nil {
			fmt.Printf("Failed to list clothes: %v\n", err)
			os.Exit(1)
		}

		if len(pieces) == 0 {
			fmt.Println("There are no clothes in the database")
			os.Exit(0)
		}

		fmt.Printf("Found %d clothes:\n", len(pieces))
		for _, p := range pieces {
			wears, _ := clothes.GetWears(db, p.ID)
			if p.ID == p.Description {
				fmt.Printf(" - %s (wears=%d)\n", p.ID, wears)
			} else {
				fmt.Printf(" - %s (%s, wears=%d)\n", p.ID, p.Description, wears)
			}
		}
	},
}
