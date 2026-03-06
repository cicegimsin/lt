package cmd

import (
	"github.com/spf13/cobra"
	"github.com/cicegimsin/lt/internal/ui"
	"github.com/cicegimsin/lt/pkg/update"
)

var updateCmd = &cobra.Command{
	Use:     "güncelle",
	Aliases: []string{"update"},
	Short:   "Sistem ve AUR paketlerini güncelle",
	Run: func(cmd *cobra.Command, args []string) {
		ui.Info("Paketler kontrol ediliyor...")
		
		updater := update.New(cfg, tr)
		if err := updater.Update(); err != nil {
			ui.Error("Güncelleme başarısız: %v", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
