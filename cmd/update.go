package cmd

import (
	"github.com/spf13/cobra"
	"github.com/cicegimsin/lt/internal/ui"
	"github.com/cicegimsin/lt/internal/universal"
)

var updateCmd = &cobra.Command{
	Use:     "güncelle",
	Aliases: []string{"update"},
	Short:   "Sistem paketlerini güncelle (evrensel)",
	Run: func(cmd *cobra.Command, args []string) {
		um, err := universal.NewUniversalManager()
		if err != nil {
			ui.Error("Sistem algılanamadı: %v", err)
			return
		}
		
		ui.Banner("SİSTEM GÜNCELLEMESİ")
		ui.Info("Sistem: %s", um.GetSystemInfo())
		ui.Info("Paketler güncelleniyor...")
		
		if err := um.Update(cfg.NoConfirm); err != nil {
			ui.Error("Güncelleme başarısız: %v", err)
			return
		}
		
		ui.Success("Sistem başarıyla güncellendi!")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
