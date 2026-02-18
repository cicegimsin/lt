package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/cicegimsin/lt-aur-helper/internal/ui"
)

var cleanCmd = &cobra.Command{
	Use:     "temizle",
	Aliases: []string{"clean"},
	Short:   "Eski önbellek dosyalarını temizle",
	Run: func(cmd *cobra.Command, args []string) {
		ui.Info("Önbellek temizleniyor...")
		
		if cfg == nil || cfg.CacheDir == "" {
			ui.Error("Yapılandırma yüklenemedi")
			return
		}
		
		if _, err := os.Stat(cfg.CacheDir); os.IsNotExist(err) {
			ui.Warning("Önbellek dizini bulunamadı")
			return
		}
		
		entries, err := os.ReadDir(cfg.CacheDir)
		if err != nil {
			ui.Error("Önbellek okunamadı: %v", err)
			return
		}
		
		count := 0
		for _, entry := range entries {
			if entry.IsDir() {
				path := filepath.Join(cfg.CacheDir, entry.Name())
				if err := os.RemoveAll(path); err == nil {
					count++
				}
			}
		}
		
		ui.Success("%d dizin temizlendi", count)
		fmt.Printf("Önbellek konumu: %s\n", cfg.CacheDir)
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
