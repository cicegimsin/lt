package search

import (
	"fmt"
	"time"

	"github.com/cicegimsin/lt/internal/aur"
	"github.com/cicegimsin/lt/internal/i18n"
	"github.com/cicegimsin/lt/internal/ui"
)

func Search(query string) ([]aur.AURPackage, error) {
	client := aur.NewClient()
	return client.Search(query)
}

func DisplayResults(results []aur.AURPackage, tr *i18n.Translator) {
	for _, pkg := range results {
		fmt.Printf("\n%s %s ★ %d\n",
			ui.Bold(pkg.Name),
			pkg.Version,
			pkg.NumVotes,
		)
		
		if pkg.Description != "" {
			fmt.Printf("   %s\n", pkg.Description)
		}
		
		lastMod := time.Unix(pkg.LastModified, 0)
		fmt.Printf("   Güncelleme: %s\n", lastMod.Format("2006-01-02"))
	}
}
