package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/cicegimsin/lt/internal/ui"
	"github.com/cicegimsin/lt/internal/universal"
)

var (
	noConfirm bool
)

var installCmd = &cobra.Command{
	Use:     "kur [paket]",
	Aliases: []string{"install"},
	Short:   "Paket kur (evrensel)",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pkgName := args[0]
		
		um, err := universal.NewUniversalManager()
		if err != nil {
			ui.Error("Sistem tespit edilemedi: %v", err)
			return
		}
		
		ui.Banner("PAKET KURULUMU")
		ui.Info("Sistem: %s", um.GetSystemInfo())
		ui.Info("Paket kuruluyor: %s", pkgName)
		
		if noConfirm {
			cfg.NoConfirm = true
		}
		
		if !cfg.NoConfirm {
			if !askConfirmation(fmt.Sprintf("'%s' paketini kurmak istiyor musunuz?", pkgName)) {
				ui.Warning("Kurulum iptal edildi")
				return
			}
		}
		
		if err := um.Install([]string{pkgName}, cfg.NoConfirm); err != nil {
			ui.Error("Kurulum başarısız: %v", err)
			return
		}
		
		ui.Success("Paket başarıyla kuruldu: %s", pkgName)
	},
}

func askConfirmation(message string) bool {
	fmt.Printf("%s [E/h] (varsayılan: E): ", message)
	
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}
	
	response = strings.ToLower(strings.TrimSpace(response))
	return response == "" || response == "e" || response == "evet" || response == "y" || response == "yes"
}

func init() {
	installCmd.Flags().BoolVar(&noConfirm, "noconfirm", false, "onay istemeden kur")
	rootCmd.AddCommand(installCmd)
}
