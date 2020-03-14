package cmd

import (
	"errors"
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteMovieCmd)

	deleteMovieCmd.Flags().StringVar(&listName, "list", "default", "enter list name")
	deleteMovieCmd.MarkFlagRequired("list")
	addCmd.Flags().StringVar(&listName, "list", "default", "enter list name")
	addCmd.MarkFlagRequired("list")
	addCmd.Flags().Float32VarP(&rate, "rate", "r", 1, "type like add --rate = 5.4")
}

var listName string
var rate float32

var addCmd = &cobra.Command{
	Use:   "add [movie]",
	Short: "adds movie to list",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a movie name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		movieName := strings.Join(args, " ")
		err := Dbs.SaveMovie(movieName, listName, rate)
		if err != nil {
			log.Fatal(err)
		}
		color.Green("  %s added to %s", movieName, listName)
	},
}

var deleteMovieCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete movies from lists",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		movieName := strings.Join(args, " ")
		err := Dbs.DeleteMovie(movieName, listName)
		if err != nil {
			log.Fatal(err)
		}
		color.Green("  %s deleted from %s", movieName, listName)
	},
}
