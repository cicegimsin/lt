package search

import (
	"fmt"
	"sort"
	"strings"
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
	repoCount := 0
	aurCount := 0
	
	for _, result := range results {
		if result.IsAUR {
			aurCount++
		} else {
			repoCount++
		}
	}
	
	if repoCount > 0 {
		ui.Header(fmt.Sprintf("Resmi Repolar (%d paket)", repoCount))
		for _, result := range results {
			if !result.IsAUR {
				fmt.Printf("%s/%s %s\n",
					ui.Repository(result.Repository),
					ui.Bold(result.Name),
					ui.Highlight(result.Version),
				)
				if result.Description != "" {
					fmt.Printf("    %s\n", result.Description)
				}
				fmt.Println()
			}
		}
	}
	
	if aurCount > 0 {
		ui.Header(fmt.Sprintf("AUR (%d paket)", aurCount))
		for _, result := range results {
			if result.IsAUR {
				votes := ""
				if result.Votes > 0 {
					votes = fmt.Sprintf(" ★ %d", result.Votes)
				}
				
				fmt.Printf("%s/%s %s%s\n",
					ui.Repository("aur"),
					ui.Bold(result.Name),
					ui.Highlight(result.Version),
					votes,
				)
				
				if result.Description != "" {
					desc := result.Description
					if len(desc) > 80 {
						desc = desc[:77] + "..."
					}
					fmt.Printf("    %s\n", desc)
				}
				
				if !result.LastMod.IsZero() {
					fmt.Printf("    Son güncelleme: %s\n", 
						result.LastMod.Format("02.01.2006"))
				}
				fmt.Println()
			}
		}
	}
	
	if len(results) == 0 {
		ui.Warning("Hiçbir paket bulunamadı")
	} else {
		ui.Separator()
		ui.Info(fmt.Sprintf("Toplam %d paket bulundu", len(results)))
	}
}
