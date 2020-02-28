package cmd

import (
	"github.com/furkanmavili/govie/pkg/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(suggestCmd)
	suggestCmd.Flags().StringVarP(&genre, "genre", "g", "", "enter type of search. eg: -t=movie or -t=series")
}

var rating int
var genre string
var suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "random movie suggester",
	Run: func(cmd *cobra.Command, args []string) {
		api.FilterGenre(genre)
	},
}
