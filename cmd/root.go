package cmd

import (
	"fmt"
	"os"

	"github.com/furkanmavili/govie/pkg/database"
	"github.com/furkanmavili/govie/pkg/database/sqlite"
	"github.com/spf13/cobra"
)

// Dbs interface for all commands
var Dbs database.Service

var rootCmd = &cobra.Command{
	Use:   "govie [command]",
	Short: "Govie is searcher of movie/tv shows",
	Args:  cobra.MinimumNArgs(1),
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if Dbs != nil {
			_ = Dbs.Close()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

//Execute for cobra
func Execute() {
	//api.SaveGenres()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	var err error
	Dbs, err = sqlite.New()
	if err != nil {
		fmt.Println(err)
	}
}
