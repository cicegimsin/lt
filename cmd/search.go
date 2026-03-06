package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/cicegimsin/lt/internal/ui"
	"github.com/cicegimsin/lt/pkg/search"
)

var searchCmd = &cobra.Command{
	Use:     "ara [paket]",
	Aliases: []string{"search"},
	Short:   "Resmi repolar ve AUR'da paket ara",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		
		ui.Info("Paket aranıyor: %s", query)
		
		results, err := search.Search(query, cfg.PacmanPath, cfg.SudoPath)
		if err != nil {
			ui.Error("Arama başarısız: %v", err)
			return
		}
		
		if len(results) == 0 {
			ui.Warning("Paket bulunamadı")
			return
		}
		
		search.DisplayResults(results, tr)
		fmt.Println()
		ui.Info("Kurulum için: lt kur <paket-adı>")
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
