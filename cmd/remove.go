package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/cicegimsin/lt/internal/ui"
	"github.com/cicegimsin/lt/internal/universal"
)

var removeCmd = &cobra.Command{
	Use:     "kaldır [paket]",
	Aliases: []string{"remove"},
	Short:   "Paketi kaldır (evrensel)",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pkgName := args[0]
		
		um, err := universal.NewUniversalManager()
		if err != nil {
			ui.Error("Sistem tespit edilemedi: %v", err)
			return
		}
		
		ui.Banner("PAKET KALDIRMA")
		ui.Info("Sistem: %s", um.GetSystemInfo())
		
		if !um.IsInstalled(pkgName) {
			ui.Error("'%s' paketi kurulu değil", pkgName)
			return
		}
		
		ui.Box("UYARI", fmt.Sprintf("'%s' paketi kaldırılacak", pkgName))
		
		if !cfg.NoConfirm {
			fmt.Printf("\nDevam etmek istiyor musunuz? [E/h] (varsayılan: h): ")
			reader := bufio.NewReader(os.Stdin)
			response, _ := reader.ReadString('\n')
			response = strings.TrimSpace(strings.ToLower(response))
			
			if response != "e" && response != "evet" && response != "y" && response != "yes" {
				ui.Box("İPTAL", "Paket kaldırma işlemi iptal edildi")
				return
			}
		}
		
		ui.Info("'%s' kaldırılıyor...", pkgName)
		
		if err := um.Remove([]string{pkgName}, cfg.NoConfirm); err != nil {
			ui.Error("Paket kaldırılamadı: %v", err)
			return
		}
		
		ui.Success("Paket başarıyla kaldırıldı: %s", pkgName)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
