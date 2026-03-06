package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/cicegimsin/lt/internal/pacman"
	"github.com/cicegimsin/lt/internal/ui"
)

var listCmd = &cobra.Command{
	Use:     "liste",
	Aliases: []string{"list"},
	Short:   "Kurulu paketleri listele",
	Run: func(cmd *cobra.Command, args []string) {
		pacmanClient := pacman.NewClient(cfg.PacmanPath, cfg.SudoPath)
		
		ui.Header("KURULU PAKETLER")
		ui.Info("Paketler listeleniyor...")
		
		packages, err := pacmanClient.GetForeignPackages()
		if err != nil {
			ui.Error("Paket listesi alınamadı: %v", err)
			return
		}
		
		if len(packages) == 0 {
			ui.Box("BİLGİ", "Kurulu paket bulunamadı")
			return
		}
		
		var items []string
		for _, pkg := range packages {
			items = append(items, fmt.Sprintf("%s %s", 
				ui.Bold(pkg.Name), 
				ui.Highlight(pkg.Version)))
		}
		
		ui.CategoryBox("Kurulu Paketler", fmt.Sprintf("%d paket", len(packages)), items)
		
		ui.Separator()
		ui.Box("ÖZET", fmt.Sprintf("Toplam %d paket kurulu", len(packages)))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
