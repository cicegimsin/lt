package cmd

import (
	"github.com/spf13/cobra"
	"github.com/cicegimsin/lt-aur-helper/internal/ui"
	"github.com/cicegimsin/lt-aur-helper/pkg/install"
)

var installCmd = &cobra.Command{
	Use:     "kur [paket]",
	Aliases: []string{"install"},
	Short:   "Paket kur",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pkgName := args[0]
		
		ui.Info(tr.Get("installing"), pkgName)
		
		installer := install.New(cfg, tr)
		if err := installer.Install(pkgName); err != nil {
			ui.Error(tr.Get("install_failed"), err)
			return
		}
		
		ui.Success(tr.Get("install_complete"), pkgName)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
