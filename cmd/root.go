package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/furkanmavili/govie/pkg/database"
	"github.com/furkanmavili/govie/pkg/database/sqlite"
	"github.com/spf13/cobra"
)

// Dbs interface for all commands
var Dbs database.Service

var rootCmd = &cobra.Command{
	Use:   "app [command]",
	Short: "Govie is searcher of movie/tv shows",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if Dbs == nil {
			log.Fatalf("Error: dbs interface is not defined.\n")
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if Dbs != nil {
			_ = Dbs.Close()
			Dbs = nil
		}
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
