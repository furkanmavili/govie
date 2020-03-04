package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(createListCmd)
	rootCmd.AddCommand(deleteListCmd)
	rootCmd.AddCommand(exportListCmd)
	listCmd.AddCommand(listAllCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list command for all list operations",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a list name")
		}
		if len(args) > 1 {
			return errors.New("list name can not be more than 1 word")
		}

		return fmt.Errorf("invalid list name")
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var createListCmd = &cobra.Command{
	Use:   "create [list name]",
	Short: "create list",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a list name")
		}
		if len(args) > 1 {
			return errors.New("list name can not be more than 1 word")
		}
		if !Dbs.IsValid(args[0]) {
			return errors.New("list already exist")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// err := Dbs.CreateList(args[0])
		err := Dbs.CreateTable(args[0])
		if err != nil {
			fmt.Println(err)
		}
		color.Green("  %s has created.", args[0])
	},
}

var deleteListCmd = &cobra.Command{
	Use:   "delete [list name]",
	Short: "deletes list",
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
		err := Dbs.DeleteList(args[0])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("List deleted:", args[0])
	},
}

// List exportlamayı eklemeyi unutma,daha sonra
var exportListCmd = &cobra.Command{
	Use:   "export [list name]",
	Short: "export list as .csv file",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a list name")
		}
		if len(args) > 1 {
			return errors.New("list name can not be more than 1 word")
		}

		return fmt.Errorf("invalid list name")
	},
	Run: func(cmd *cobra.Command, args []string) {
		new := strings.Join(args, "_")
		fmt.Printf("List has exported: %s.csv\n", new)
	},
}

var listAllCmd = &cobra.Command{
	Use:   "all",
	Short: "shows all lists",
	Run: func(cmd *cobra.Command, args []string) {
		err := Dbs.ShowListsAll()
		if err != nil {
			fmt.Println(err)
		}
	},
}
