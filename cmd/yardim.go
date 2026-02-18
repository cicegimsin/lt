package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	// completion komutunu gizle
	for _, cmd := range rootCmd.Commands() {
		if cmd.Name() == "completion" {
			cmd.Hidden = true
		}
	}
	
	// help komutunu Türkçeleştir
	helpCmd := &cobra.Command{
		Use:   "yardim [komut]",
		Short: "Herhangi bir komut hakkında yardım",
		Long: `Herhangi bir komut hakkında yardım ve kullanım bilgisi gösterir.

Örnek:
  lt yardim ara
  lt yardim kur`,
		Run: func(c *cobra.Command, args []string) {
			cmd, _, e := c.Root().Find(args)
			if cmd == nil || e != nil {
				c.Printf("Bilinmeyen komut: %q\n", args)
				return
			}
			cmd.Help()
		},
	}
	
	rootCmd.AddCommand(helpCmd)
	
	// Orijinal help komutunu gizle
	rootCmd.SetHelpCommand(helpCmd)
}
