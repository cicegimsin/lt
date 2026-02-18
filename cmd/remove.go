package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/cicegimsin/lt/internal/ui"
)

var removeCmd = &cobra.Command{
	Use:     "kaldır [paket]",
	Aliases: []string{"remove"},
	Short:   "Paketi kaldır",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pkgName := args[0]
		
		// Paket kurulu mu kontrol et
		checkCmd := exec.Command("pacman", "-Q", pkgName)
		if err := checkCmd.Run(); err != nil {
			ui.Error("'%s' paketi kurulu değil", pkgName)
			return
		}
		
		// Onay iste
		fmt.Printf("\n'%s' paketini kaldırmak istediğinize emin misiniz? [E/h]: ", ui.Bold(pkgName))
		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))
		
		if response != "e" && response != "evet" && response != "y" && response != "yes" {
			ui.Info("İşlem iptal edildi")
			return
		}
		
		ui.Info("'%s' kaldırılıyor...", pkgName)
		
		cmdExec := exec.Command("sudo", "pacman", "-Rns", pkgName)
		cmdExec.Stdout = os.Stdout
		cmdExec.Stderr = os.Stderr
		cmdExec.Stdin = os.Stdin
		
		if err := cmdExec.Run(); err != nil {
			ui.Error("Paket kaldırılamadı: %v", err)
			return
		}
		
		ui.Success("Paket kaldırıldı: %s", pkgName)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
