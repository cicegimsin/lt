package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/cicegimsin/lt/internal/ui"
	"github.com/cicegimsin/lt/internal/universal"
)

var searchCmd = &cobra.Command{
	Use:     "ara [paket]",
	Aliases: []string{"search"},
	Short:   "Paket ara (tüm kaynaklardan)",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		
		um, err := universal.NewUniversalManager()
		if err != nil {
			ui.Error("Sistem tespit edilemedi: %v", err)
			return
		}
		
		ui.Banner("PAKET ARAMA")
		ui.Info("Sistem: %s", um.GetSystemInfo())
		ui.Info("Paket aranıyor: %s", query)
		
		results, err := um.Search(query)
		if err != nil {
			ui.Error("Arama başarısız: %v", err)
			return
		}
		
		if len(results) == 0 {
			ui.Warning("Hiçbir paket bulunamadı")
			return
		}
		
		displayUniversalResults(results)
		
		ui.Separator()
		ui.Box("ÖZET", fmt.Sprintf("Toplam %d paket bulundu\nKurulum için: lt kur <paket-adı>", len(results)))
	},
}

func displayUniversalResults(results []universal.Package) {
	officialCount := 0
	communityCount := 0
	otherCount := 0
	
	for _, result := range results {
		switch result.Source {
		case "official":
			officialCount++
		case "aur", "community":
			communityCount++
		default:
			otherCount++
		}
	}
	
	if officialCount > 0 {
		var officialItems []string
		for _, result := range results {
			if result.Source == "official" {
				item := fmt.Sprintf("%s %s", 
					ui.Bold(result.Name),
					ui.Highlight(result.Version))
				
				if result.Repository != "" {
					item = fmt.Sprintf("%s/%s", 
						ui.Repository(result.Repository), item)
				}
				
				if result.Description != "" {
					desc := result.Description
					if len(desc) > 60 {
						desc = desc[:57] + "..."
					}
					item += "\n" + desc
				}
				officialItems = append(officialItems, item)
			}
		}
		ui.CategoryBox("Resmi Paketler", fmt.Sprintf("%d paket", officialCount), officialItems)
		fmt.Println()
	}
	
	if communityCount > 0 {
		var communityItems []string
		for _, result := range results {
			if result.Source == "aur" || result.Source == "community" {
				item := fmt.Sprintf("%s/%s %s",
					ui.Repository("topluluk"),
					ui.Bold(result.Name),
					ui.Highlight(result.Version))
				
				if result.Description != "" {
					desc := result.Description
					if len(desc) > 60 {
						desc = desc[:57] + "..."
					}
					item += "\n" + desc
				}
				
				communityItems = append(communityItems, item)
			}
		}
		ui.CategoryBox("Topluluk Paketleri", fmt.Sprintf("%d paket", communityCount), communityItems)
		fmt.Println()
	}
	
	if otherCount > 0 {
		var otherItems []string
		for _, result := range results {
			if result.Source != "official" && result.Source != "aur" && result.Source != "community" {
				item := fmt.Sprintf("%s %s", 
					ui.Bold(result.Name),
					ui.Highlight(result.Version))
				
				if result.Description != "" {
					desc := result.Description
					if len(desc) > 60 {
						desc = desc[:57] + "..."
					}
					item += "\n" + desc
				}
				otherItems = append(otherItems, item)
			}
		}
		ui.CategoryBox("Diğer Kaynaklar", fmt.Sprintf("%d paket", otherCount), otherItems)
	}
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
