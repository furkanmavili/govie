package cmd

import (
	"fmt"

	"github.com/furkanmavili/govie/pkg/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(suggestCmd)
	suggestCmd.Flags().StringVarP(&genre, "genre", "g", "", "enter type of search. eg: -t=movie or -t=series")
}

var genre string
var suggestCmd = &cobra.Command{
	Use:   "suggest [flag]",
	Short: "random movie suggester",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(genre)
		api.FilterGenre(genre)
	},
}
