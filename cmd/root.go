package cmd

import (
	"fmt"
	"os"

	"github.com/furkanmavili/govie/pkg/api"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app [command]",
	Short: "Govie is searcher of movie/tv shows",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("to search movies or series use 'search' command.\n" +
			"to get suggestion use 'suggest' command")
	},
}

//Execute for cobra
func Execute() {
	api.SaveGenres()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
