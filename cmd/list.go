package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listAllCmd)
	listCmd.AddCommand(createListCmd)
	listCmd.AddCommand(deleteListCmd)
	listCmd.AddCommand(exportListCmd)
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
		err := Dbs.CreateList(args[0])
		if err != nil {
			fmt.Println("hata")
		}
		fmt.Println("List oluşturuldu:", args[0])
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
		fmt.Println("List silindi:", args[0])
	},
}

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
		fmt.Printf("List export edildi: %s.csv\n", new)
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
