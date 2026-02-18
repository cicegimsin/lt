package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/clicksama/lt/internal/aur"
	"github.com/clicksama/lt/internal/ui"
)

var infoCmd = &cobra.Command{
	Use:     "bilgi [paket]",
	Aliases: []string{"info"},
	Short:   "Paket detaylarını göster",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pkgName := args[0]
		
		client := aur.NewClient()
		pkg, err := client.Info(pkgName)
		if err != nil {
			ui.Error("Paket bilgisi alınamadı: %v", err)
			return
		}
		
		fmt.Printf("\n%s %s\n", ui.Bold(pkg.Name), pkg.Version)
		fmt.Printf("Açıklama: %s\n", pkg.Description)
		fmt.Printf("Adres: %s\n", pkg.URL)
		fmt.Printf("Oy Sayısı: %d\n", pkg.NumVotes)
		fmt.Printf("Popülerlik: %.2f\n", pkg.Popularity)
		fmt.Printf("Bakımcı: %s\n", pkg.Maintainer)
		
		lastMod := time.Unix(pkg.LastModified, 0)
		fmt.Printf("Son Güncelleme: %s\n\n", lastMod.Format("2006-01-02 15:04"))
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
