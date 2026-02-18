package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/cicegimsin/lt/internal/aur"
	"github.com/cicegimsin/lt/internal/ui"
)

var downloadCmd = &cobra.Command{
	Use:     "indir [paket]",
	Aliases: []string{"download"},
	Short:   "Sadece kaynak kodunu indir",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pkgName := args[0]
		
		client := aur.NewClient()
		pkg, err := client.Info(pkgName)
		if err != nil {
			ui.Error("Paket bilgisi alınamadı: %v", err)
			return
		}
		
		buildDir := filepath.Join(cfg.CacheDir, pkg.Name)
		if err := os.MkdirAll(buildDir, 0755); err != nil {
			ui.Error("Dizin oluşturulamadı: %v", err)
			return
		}
		
		ui.Info("'%s' indiriliyor...", pkg.Name)
		
		url := fmt.Sprintf("https://aur.archlinux.org/%s.git", pkg.Name)
		
		if _, err := os.Stat(buildDir); err == nil {
			os.RemoveAll(buildDir)
		}
		
		cmdExec := exec.Command("git", "clone", url, buildDir)
		cmdExec.Stdout = os.Stdout
		cmdExec.Stderr = os.Stderr
		
		if err := cmdExec.Run(); err != nil {
			ui.Error("İndirme başarısız: %v", err)
			return
		}
		
		ui.Success("İndirildi: %s", buildDir)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
