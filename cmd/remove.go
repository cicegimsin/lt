package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/clicksama/lt/internal/ui"
)

var removeCmd = &cobra.Command{
	Use:     "kaldır [paket]",
	Aliases: []string{"remove"},
	Short:   "Paketi kaldır",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pkgName := args[0]
		
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
