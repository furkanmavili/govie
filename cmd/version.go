package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version of goovie",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Goovie movie/tv show searcher v1.0")
	},
}
