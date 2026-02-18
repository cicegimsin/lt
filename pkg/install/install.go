package install

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/clicksama/lt/internal/aur"
	"github.com/clicksama/lt/internal/config"
	"github.com/clicksama/lt/internal/i18n"
	"github.com/clicksama/lt/internal/ui"
)

type Installer struct {
	cfg    *config.Config
	tr     *i18n.Translator
	client *aur.Client
}

func New(cfg *config.Config, tr *i18n.Translator) *Installer {
	return &Installer{
		cfg:    cfg,
		tr:     tr,
		client: aur.NewClient(),
	}
}

func (i *Installer) Install(pkgName string) error {
	pkg, err := i.client.Info(pkgName)
	if err != nil {
		return fmt.Errorf("paket bilgisi alınamadı: %w", err)
	}
	
	ui.Success("Bağımlılıklar kontrol ediliyor...")
	
	buildDir := filepath.Join(i.cfg.CacheDir, pkg.Name)
	if err := os.MkdirAll(buildDir, 0755); err != nil {
		return err
	}
	
	ui.Success("Kaynak indiriliyor...")
	if err := i.cloneRepo(pkg.Name, buildDir); err != nil {
		return err
	}
	
	ui.Success("Paket derleniyor...")
	if err := i.buildPackage(buildDir); err != nil {
		return err
	}
	
	return nil
}

func (i *Installer) cloneRepo(pkgName, dest string) error {
	url := fmt.Sprintf("https://aur.archlinux.org/%s.git", pkgName)
	
	if _, err := os.Stat(dest); err == nil {
		os.RemoveAll(dest)
	}
	
	cmd := exec.Command("git", "clone", url, dest)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func (i *Installer) buildPackage(buildDir string) error {
	cmd := exec.Command("makepkg", "-si", "--noconfirm")
	cmd.Dir = buildDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), "MAKEFLAGS="+i.cfg.MakeFlags)
	
	return cmd.Run()
}
