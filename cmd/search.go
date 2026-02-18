package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/cicegimsin/lt-aur-helper/internal/ui"
	"github.com/cicegimsin/lt-aur-helper/pkg/search"
)

var searchCmd = &cobra.Command{
	Use:     "ara [paket]",
	Aliases: []string{"search"},
	Short:   "AUR'da paket ara",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		
		ui.Info(tr.Get("searching"), query)
		
		results, err := search.Search(query)
		if err != nil {
			ui.Error(tr.Get("search_failed"), err)
			return
		}
		
		if len(results) == 0 {
			ui.Warning(tr.Get("no_results"))
			return
		}
		
		search.DisplayResults(results, tr)
		fmt.Println()
		ui.Info(tr.Get("install_hint"))
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
