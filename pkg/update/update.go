package update

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cicegimsin/lt/internal/aur"
	"github.com/cicegimsin/lt/internal/config"
	"github.com/cicegimsin/lt/internal/i18n"
	"github.com/cicegimsin/lt/internal/ui"
	"github.com/cicegimsin/lt/pkg/install"
)

type Updater struct {
	cfg    *config.Config
	tr     *i18n.Translator
	client *aur.Client
}

type UpdateInfo struct {
	Name       string
	LocalVer   string
	RemoteVer  string
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
	
	var updates []UpdateInfo
	for _, pkg := range installed {
		aurPkg, err := u.client.Info(pkg)
		if err != nil {
			continue
		}
		
		localVer := u.getLocalVersion(pkg)
		if localVer != aurPkg.Version {
			updates = append(updates, UpdateInfo{
				Name:      pkg,
				LocalVer:  localVer,
				RemoteVer: aurPkg.Version,
			})
		}
	}
	
	if len(updates) == 0 {
		ui.Success("Tüm paketler güncel")
		return nil
	}
	
	fmt.Printf("\n[+] %d paket güncellenebilir:\n", len(updates))
	for _, upd := range updates {
		fmt.Printf("  - %s (%s -> %s)\n", ui.Bold(upd.Name), upd.LocalVer, upd.RemoteVer)
	}
	
	fmt.Print("\nGüncellemek istiyor musunuz? [E/h] (varsayılan: E): ")
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		ui.Error("Girdi okunamadı: %v", err)
		return err
	}
	response = strings.TrimSpace(strings.ToLower(response))
	
	// Boş girdi = varsayılan evet
	if response == "" {
		response = "e"
	}
	
	// Sadece açıkça "hayır" derse iptal et
	if response == "h" || response == "hayir" || response == "n" || response == "no" {
		ui.Info("Güncelleme iptal edildi")
		return nil
	}
	
	fmt.Println()
	installer := install.New(u.cfg, u.tr)
	
	for i, upd := range updates {
		fmt.Printf("[%d/%d] %s güncelleniyor...\n", i+1, len(updates), upd.Name)
		
		if err := installer.Install(upd.Name); err != nil {
			ui.Error("%s güncellenemedi: %v", upd.Name, err)
			continue
		}
		
		ui.Success("%s güncellendi (%s)", upd.Name, upd.RemoteVer)
	}
	
	fmt.Println()
	ui.Success("Güncelleme tamamlandı!")
	
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
