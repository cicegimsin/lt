package cmd

import (
	"github.com/spf13/cobra"
	"github.com/cicegimsin/lt-aur-helper/internal/ui"
	"github.com/cicegimsin/lt-aur-helper/pkg/update"
)

var updateCmd = &cobra.Command{
	Use:     "güncelle",
	Aliases: []string{"update"},
	Short:   "Tüm AUR paketlerini güncelle",
	Run: func(cmd *cobra.Command, args []string) {
		ui.Info(tr.Get("checking_updates"))
		
		updater := update.New(cfg, tr)
		if err := updater.Update(); err != nil {
			ui.Error(tr.Get("update_failed"), err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
