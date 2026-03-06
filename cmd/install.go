package cmd

import (
	"github.com/spf13/cobra"
	"github.com/cicegimsin/lt/internal/ui"
	"github.com/cicegimsin/lt/pkg/install"
)

var (
	noConfirm bool
)

var installCmd = &cobra.Command{
	Use:     "kur [paket]",
	Aliases: []string{"install"},
	Short:   "Paket kur",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pkgName := args[0]
		
		// NoConfirm flag'ini config'e uygula
		if noConfirm {
			cfg.NoConfirm = true
		}
		
		ui.Info("'%s' kuruluyor...", pkgName)
		
		installer := install.New(cfg, tr)
		if err := installer.Install(pkgName); err != nil {
			ui.Error("Kurulum başarısız: %v", err)
			return
		}
		
		ui.Success("Kurulum tamamlandı: %s", pkgName)
	},
}

func init() {
	installCmd.Flags().BoolVar(&noConfirm, "noconfirm", false, "onay istemeden kur")
	rootCmd.AddCommand(installCmd)
}
