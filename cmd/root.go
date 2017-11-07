package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(listCmd, addCmd, wearCmd, washCmd)
}

var RootCmd = &cobra.Command{
	Use: "diwia",
}
