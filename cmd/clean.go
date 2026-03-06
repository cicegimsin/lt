package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/cicegimsin/lt/internal/ui"
)

var cleanCmd = &cobra.Command{
	Use:     "temizle",
	Aliases: []string{"clean"},
	Short:   "Eski önbellek dosyalarını temizle",
	Run: func(cmd *cobra.Command, args []string) {
		ui.Header("ÖNBELLEK TEMİZLİĞİ")
		
		if cfg == nil || cfg.CacheDir == "" {
			ui.Error("Yapılandırma yüklenemedi")
			return
		}
		
		ui.Info("Önbellek dizini kontrol ediliyor...")
		
		if _, err := os.Stat(cfg.CacheDir); os.IsNotExist(err) {
			ui.Box("BİLGİ", "Önbellek dizini bulunamadı\nKonum: "+cfg.CacheDir)
			return
		}
		
		entries, err := os.ReadDir(cfg.CacheDir)
		if err != nil {
			ui.Error("Önbellek okunamadı: %v", err)
			return
		}
		
		if len(entries) == 0 {
			ui.Box("BİLGİ", "Önbellek zaten temiz")
			return
		}
		
		ui.Info("Önbellek temizleniyor...")
		
		count := 0
		var cleanedDirs []string
		
		for _, entry := range entries {
			if entry.IsDir() {
				path := filepath.Join(cfg.CacheDir, entry.Name())
				if err := os.RemoveAll(path); err == nil {
					count++
					cleanedDirs = append(cleanedDirs, entry.Name())
				}
			}
		}
		
		if count > 0 {
			ui.Success("Önbellek temizlendi")
			
			if len(cleanedDirs) <= 10 {
				ui.Box("TEMİZLENEN DİZİNLER", fmt.Sprintf("Toplam: %d dizin\n\n%s", 
					count, 
					fmt.Sprintf("• %s", fmt.Sprintf("%s\n• ", cleanedDirs))))
			} else {
				ui.Box("ÖZET", fmt.Sprintf("Toplam %d dizin temizlendi\nKonum: %s", count, cfg.CacheDir))
			}
		} else {
			ui.Warning("Hiçbir dizin temizlenemedi")
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
