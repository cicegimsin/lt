package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/cicegimsin/lt/internal/aur"
	"github.com/cicegimsin/lt/internal/pacman"
	"github.com/cicegimsin/lt/internal/ui"
)

var infoCmd = &cobra.Command{
	Use:     "bilgi [paket]",
	Aliases: []string{"info"},
	Short:   "Paket detaylarını göster",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pkgName := args[0]
		
		pacmanClient := pacman.NewClient(cfg.PacmanPath, cfg.SudoPath)
		
		// Önce resmi repolardan ara
		repoPackages, err := pacmanClient.Search(pkgName)
		if err == nil {
			for _, pkg := range repoPackages {
				if pkg.Name == pkgName {
					showRepoPackageInfo(pkg)
					return
				}
			}
		}
		
		// AUR'dan ara
		aurClient := aur.NewClient()
		pkg, err := aurClient.Info(pkgName)
		if err != nil {
			ui.Error("Paket bulunamadı: %v", err)
			return
		}
		
		showAURPackageInfo(pkg, pacmanClient)
	},
}

func showRepoPackageInfo(pkg pacman.Package) {
	ui.Header("PAKET BİLGİLERİ")
	
	info := fmt.Sprintf("Ad: %s\nSürüm: %s\nRepo: %s", 
		pkg.Name, 
		pkg.Version, 
		strings.ToUpper(pkg.Repository))
	
	if pkg.Description != "" {
		info += "\nAçıklama: " + pkg.Description
	}
	
	ui.Box("RESMİ REPO PAKETİ", info)
}

func showAURPackageInfo(pkg *aur.AURPackage, pacmanClient *pacman.Client) {
	ui.Header("PAKET BİLGİLERİ")
	
	status := "Kurulu değil"
	if pacmanClient.IsInstalled(pkg.Name) {
		status = ui.Highlight("Kurulu")
	}
	
	lastMod := time.Unix(pkg.LastModified, 0)
	
	info := fmt.Sprintf("Ad: %s\nSürüm: %s\nDurum: %s\nOy Sayısı: %d\nPopülerlik: %.2f\nBakımcı: %s\nSon Güncelleme: %s", 
		pkg.Name,
		pkg.Version,
		status,
		pkg.NumVotes,
		pkg.Popularity,
		pkg.Maintainer,
		lastMod.Format("02.01.2006 15:04"))
	
	if pkg.Description != "" {
		info += "\nAçıklama: " + pkg.Description
	}
	
	if pkg.URL != "" {
		info += "\nWeb Sitesi: " + pkg.URL
	}
	
	ui.Box("TOPLULUK PAKETİ", info)
	
	// Bağımlılıkları göster
	if len(pkg.Depends) > 0 || len(pkg.MakeDepends) > 0 {
		ui.Header("BAĞIMLILIKLAR")
		
		if len(pkg.Depends) > 0 {
			ui.Box("Çalışma Zamanı", strings.Join(pkg.Depends, "\n"))
		}
		
		if len(pkg.MakeDepends) > 0 {
			ui.Box("Derleme Zamanı", strings.Join(pkg.MakeDepends, "\n"))
		}
		
		if len(pkg.OptDepends) > 0 {
			ui.Box("İsteğe Bağlı", strings.Join(pkg.OptDepends, "\n"))
		}
	}
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
