package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var tamamlamaCmd = &cobra.Command{
	Use:   "tamamlama [bash|zsh|fish|powershell]",
	Short: "Kabuk için otomatik tamamlama betiği oluştur",
	Long: `Belirtilen kabuk için otomatik tamamlama betiği oluşturur.

Bash için:
  lt tamamlama bash > /etc/bash_completion.d/lt

Zsh için:
  lt tamamlama zsh > "${fpath[1]}/_lt"

Fish için:
  lt tamamlama fish > ~/.config/fish/completions/lt.fish

PowerShell için:
  lt tamamlama powershell > lt.ps1
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(tamamlamaCmd)
}
