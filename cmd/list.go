package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/clicksama/lt/internal/ui"
)

var listCmd = &cobra.Command{
	Use:     "liste",
	Aliases: []string{"list"},
	Short:   "Kurulu AUR paketlerini listele",
	Run: func(cmd *cobra.Command, args []string) {
		ui.Info("Kurulu AUR paketleri:")
		
		cmdExec := exec.Command("pacman", "-Qm")
		output, err := cmdExec.Output()
		if err != nil {
			ui.Error("Paket listesi alınamadı: %v", err)
			return
		}
		
		lines := strings.Split(string(output), "\n")
		count := 0
		for _, line := range lines {
			if line == "" {
				continue
			}
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				fmt.Printf("  %s %s\n", ui.Bold(parts[0]), parts[1])
				count++
			}
		}
		
		fmt.Printf("\nToplam: %d paket\n", count)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
