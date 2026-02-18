package update

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/clicksama/lt/internal/aur"
	"github.com/clicksama/lt/internal/config"
	"github.com/clicksama/lt/internal/i18n"
	"github.com/clicksama/lt/internal/ui"
)

type Updater struct {
	cfg    *config.Config
	tr     *i18n.Translator
	client *aur.Client
}

func New(cfg *config.Config, tr *i18n.Translator) *Updater {
	return &Updater{
		cfg:    cfg,
		tr:     tr,
		client: aur.NewClient(),
	}
}

func (u *Updater) Update() error {
	installed, err := u.getInstalledAURPackages()
	if err != nil {
		return err
	}
	
	if len(installed) == 0 {
		ui.Info("Kurulu AUR paketi bulunamadı")
		return nil
	}
	
	var updates []string
	for _, pkg := range installed {
		aurPkg, err := u.client.Info(pkg)
		if err != nil {
			continue
		}
		
		localVer := u.getLocalVersion(pkg)
		if localVer != aurPkg.Version {
			updates = append(updates, fmt.Sprintf("- %s (%s -> %s)", pkg, localVer, aurPkg.Version))
		}
	}
	
	if len(updates) == 0 {
		ui.Success("Tüm paketler güncel")
		return nil
	}
	
	fmt.Printf("\n[+] %d paket güncellenebilir:\n", len(updates))
	for _, upd := range updates {
		fmt.Println(upd)
	}
	
	return nil
}

func (u *Updater) getInstalledAURPackages() ([]string, error) {
	cmd := exec.Command("pacman", "-Qm")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	
	var packages []string
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) > 0 {
			packages = append(packages, parts[0])
		}
	}
	
	return packages, nil
}

func (u *Updater) getLocalVersion(pkgName string) string {
	cmd := exec.Command("pacman", "-Q", pkgName)
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	
	parts := strings.Fields(string(output))
	if len(parts) >= 2 {
		return parts[1]
	}
	
	return ""
}
