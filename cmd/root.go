package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/clicksama/lt/internal/config"
	"github.com/clicksama/lt/internal/i18n"
)

var (
	cfgFile string
	lang    string
	cfg     *config.Config
	tr      *i18n.Translator
)

var rootCmd = &cobra.Command{
	Use:   "lt",
	Short: "Sade AUR paket yöneticisi",
	Long:  "lt - Hızlı ve sade AUR paket yöneticisi",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		cfg, err = config.Load(cfgFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Yapılandırma yüklenemedi: %v\n", err)
			os.Exit(1)
		}
		
		if lang == "" {
			lang = cfg.Language
		}
		tr = i18n.New(lang)
	},
}

func init() {
	cobra.AddTemplateFunc("tr", translateUsage)
	rootCmd.SetUsageTemplate(usageTemplate)
	rootCmd.SetHelpTemplate(helpTemplate)
}

func translateUsage(s string) string {
	translations := map[string]string{
		"Usage:":             "Kullanım:",
		"Available Commands": "Kullanılabilir Komutlar",
		"Flags:":             "Seçenekler:",
		"Global Flags:":      "Genel Seçenekler:",
		"Use":                "Kullanım",
		"for more information about a command.": "komutu hakkında daha fazla bilgi için.",
		"help for":           "yardım:",
		"version for":        "sürüm:",
		"Additional help topics:": "Ek yardım konuları:",
	}
	
	for en, tr := range translations {
		s = strings.Replace(s, en, tr, -1)
	}
	return s
}

var usageTemplate = `Kullanım:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [komut]{{end}}{{if gt (len .Aliases) 0}}

Takma Adlar:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Örnekler:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Kullanılabilir Komutlar:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasHelpSubCommands}}

Ek yardım konuları:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Daha fazla bilgi için: lt yardim [komut]{{end}}
`

var helpTemplate = `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}

{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`

func Execute() {
	initFlags()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initFlags() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "yapilandirma", "y", "", "yapılandırma dosyası")
	rootCmd.PersistentFlags().StringVarP(&lang, "dil", "d", "", "arayüz dili (tr/en)")
	rootCmd.Version = "1.0.0"
	rootCmd.SetVersionTemplate("lt sürüm {{.Version}}\n")
}
