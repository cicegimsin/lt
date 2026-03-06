package search

import (
	"fmt"
	"sort"
	"time"

	"github.com/cicegimsin/lt/internal/aur"
	"github.com/cicegimsin/lt/internal/i18n"
	"github.com/cicegimsin/lt/internal/pacman"
	"github.com/cicegimsin/lt/internal/ui"
)

type SearchResult struct {
	Name        string
	Version     string
	Description string
	Repository  string
	IsAUR       bool
	Votes       int
	Popularity  float64
	LastMod     time.Time
}

func Search(query string, pacmanPath, sudoPath string) ([]SearchResult, error) {
	var results []SearchResult
	
	pacmanClient := pacman.NewClient(pacmanPath, sudoPath)
	repoPackages, err := pacmanClient.Search(query)
	if err == nil {
		for _, pkg := range repoPackages {
			results = append(results, SearchResult{
				Name:        pkg.Name,
				Version:     pkg.Version,
				Description: pkg.Description,
				Repository:  pkg.Repository,
				IsAUR:       false,
			})
		}
	}
	
	aurClient := aur.NewClient()
	aurPackages, err := aurClient.Search(query)
	if err == nil {
		for _, pkg := range aurPackages {
			results = append(results, SearchResult{
				Name:        pkg.Name,
				Version:     pkg.Version,
				Description: pkg.Description,
				Repository:  "aur",
				IsAUR:       true,
				Votes:       pkg.NumVotes,
				Popularity:  pkg.Popularity,
				LastMod:     time.Unix(pkg.LastModified, 0),
			})
		}
	}
	
	sort.Slice(results, func(i, j int) bool {
		if results[i].IsAUR != results[j].IsAUR {
			return !results[i].IsAUR
		}
		if results[i].IsAUR {
			return results[i].Popularity > results[j].Popularity
		}
		return results[i].Name < results[j].Name
	})
	
	return results, nil
}

func DisplayResults(results []SearchResult, tr *i18n.Translator) {
	if len(results) == 0 {
		ui.Warning("Hiçbir paket bulunamadı")
		return
	}
	
	repoCount := 0
	aurCount := 0
	
	for _, result := range results {
		if result.IsAUR {
			aurCount++
		} else {
			repoCount++
		}
	}
	
	ui.Banner("PAKET ARAMA SONUÇLARI")
	
	if repoCount > 0 {
		var repoItems []string
		for _, result := range results {
			if !result.IsAUR {
				item := fmt.Sprintf("%s/%s %s", 
					ui.Repository(result.Repository),
					ui.Bold(result.Name),
					ui.Highlight(result.Version))
				
				if result.Description != "" {
					desc := result.Description
					if len(desc) > 60 {
						desc = desc[:57] + "..."
					}
					item += "\n" + desc
				}
				repoItems = append(repoItems, item)
			}
		}
		ui.CategoryBox("Resmi Repolar", fmt.Sprintf("%d paket", repoCount), repoItems)
		fmt.Println()
	}
	
	if aurCount > 0 {
		var aurItems []string
		for _, result := range results {
			if result.IsAUR {
				votes := ""
				if result.Votes > 0 {
					votes = fmt.Sprintf(" ★ %d", result.Votes)
				}
				
				item := fmt.Sprintf("%s/%s %s%s",
					ui.Repository("aur"),
					ui.Bold(result.Name),
					ui.Highlight(result.Version),
					votes)
				
				if result.Description != "" {
					desc := result.Description
					if len(desc) > 60 {
						desc = desc[:57] + "..."
					}
					item += "\n" + desc
				}
				
				if !result.LastMod.IsZero() {
					item += fmt.Sprintf("\nSon güncelleme: %s", 
						result.LastMod.Format("02.01.2006"))
				}
				
				aurItems = append(aurItems, item)
			}
		}
		ui.CategoryBox("AUR", fmt.Sprintf("%d paket", aurCount), aurItems)
	}
	
	ui.Separator()
	ui.Box("ÖZET", fmt.Sprintf("Toplam %d paket bulundu\nKurulum için: lt kur <paket-adı>", len(results)))
}
