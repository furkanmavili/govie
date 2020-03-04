package cmd

import (
	"errors"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:   "show [listName]",
	Short: "shows context in list",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a list name")
		}
		if Dbs.IsValid(args[0]) {
			return errors.New("list couln't found")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		listName := strings.Join(args, " ")
		err := Dbs.ShowList(listName)
		if err != nil {
			log.Fatal(err)
		}
	},
}
