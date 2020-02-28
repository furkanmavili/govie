package cmd

import (
	"fmt"
	"strings"

	"github.com/furkanmavili/govie/pkg/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&searchType, "type", "t", "movies", "enter type of search. eg: -t=movie or -t=series")
}

var searchType string

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "searcher",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		movie := strings.Join(args, " ")
		fmt.Printf("you wanna search %s in %s\n", movie, searchType)
		err := api.SearchMovie(movie, searchType)
		if err != nil {
			fmt.Println(err)
		}
	},
}
