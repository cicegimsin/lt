package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/cicegimsin/lt/internal/ui"
	"github.com/cicegimsin/lt/internal/universal"
)

var systemCmd = &cobra.Command{
	Use:     "sistem",
	Aliases: []string{"system", "info"},
	Short:   "Sistem bilgilerini göster",
	Run: func(cmd *cobra.Command, args []string) {
		um, err := universal.NewUniversalManager()
		if err != nil {
			ui.Error("Sistem tespit edilemedi: %v", err)
			return
		}
		
		ui.Banner("SİSTEM BİLGİLERİ")
		
		osInfo := um.OSInfo
		pm := um.PackageManager
		
		systemInfo := fmt.Sprintf("İşletim Sistemi: %s\nDağıtım: %s\nSürüm: %s\nMimari: %s", 
			osInfo.GetDisplayName(),
			osInfo.Distribution,
			osInfo.Version,
			osInfo.Type)
		
		ui.Box("İŞLETİM SİSTEMİ", systemInfo)
		
		pmInfo := fmt.Sprintf("Paket Yöneticisi: %s\nSudo Gerekli: %v\nTopluluk Desteği: %v", 
			pm.Name,
			pm.SudoNeeded,
			osInfo.SupportsCommunityRepos())
		
		ui.Box("PAKET YÖNETİCİSİ", pmInfo)
		
		commands := fmt.Sprintf("Kurulum: %s\nKaldırma: %s\nGüncelleme: %s\nArama: %s", 
			fmt.Sprintf("%s %s", pm.InstallCmd[0], pm.InstallCmd[1]),
			fmt.Sprintf("%s %s", pm.RemoveCmd[0], pm.RemoveCmd[1]),
			fmt.Sprintf("%s %s", pm.UpdateCmd[0], pm.UpdateCmd[1]),
			fmt.Sprintf("%s %s", pm.SearchCmd[0], pm.SearchCmd[1]))
		
		ui.Box("KOMUTLAR", commands)
		
		ui.Separator()
		ui.Info("lt evrensel paket yöneticisi - Tüm sistemlerde çalışır!")
	},
}

func init() {
	rootCmd.AddCommand(systemCmd)
}