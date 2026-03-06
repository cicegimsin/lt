package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/cicegimsin/lt/internal/pacman"
	"github.com/cicegimsin/lt/internal/ui"
)

var removeCmd = &cobra.Command{
	Use:     "kaldır [paket]",
	Aliases: []string{"remove"},
	Short:   "Paketi kaldır",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pkgName := args[0]
		
		// Pacman client oluştur
		pacmanClient := pacman.NewClient(cfg.PacmanPath, cfg.SudoPath)
		
		// Paket kurulu mu kontrol et
		if !pacmanClient.IsInstalled(pkgName) {
			ui.Error("'%s' paketi kurulu değil", pkgName)
			return
		}
		
		// Onay iste (eğer noconfirm değilse)
		if !cfg.NoConfirm {
			fmt.Printf("\n'%s' paketini kaldırmak istediğinize emin misiniz? [E/h]: ", ui.Bold(pkgName))
			reader := bufio.NewReader(os.Stdin)
			response, _ := reader.ReadString('\n')
			response = strings.TrimSpace(strings.ToLower(response))
			
			if response != "e" && response != "evet" && response != "y" && response != "yes" && response != "" {
				ui.Info("İşlem iptal edildi")
				return
			}
		}
		
		ui.Info("'%s' kaldırılıyor...", pkgName)
		
		// Pacman client kullanarak kaldır
		if err := pacmanClient.Remove([]string{pkgName}, cfg.NoConfirm); err != nil {
			ui.Error("Paket kaldırılamadı: %v", err)
			return
		}
		
		ui.Success("Paket kaldırıldı: %s", pkgName)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
