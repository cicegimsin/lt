package install

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/cicegimsin/lt/internal/aur"
	"github.com/cicegimsin/lt/internal/config"
	"github.com/cicegimsin/lt/internal/deps"
	"github.com/cicegimsin/lt/internal/i18n"
	"github.com/cicegimsin/lt/internal/pacman"
	"github.com/cicegimsin/lt/internal/ui"
)

type Installer struct {
	cfg           *config.Config
	tr            *i18n.Translator
	aurClient     *aur.Client
	pacmanClient  *pacman.Client
	depsResolver  *deps.Resolver
}

func New(cfg *config.Config, tr *i18n.Translator) *Installer {
	aurClient := aur.NewClient()
	pacmanClient := pacman.NewClient(cfg.PacmanPath, cfg.SudoPath)
	depsResolver := deps.NewResolver(aurClient, pacmanClient)
	
	return &Installer{
		cfg:          cfg,
		tr:           tr,
		aurClient:    aurClient,
		pacmanClient: pacmanClient,
		depsResolver: depsResolver,
	}
}

func (i *Installer) Install(pkgName string) error {
	ui.Info("Bağımlılıklar analiz ediliyor...")
	
	plan, err := i.depsResolver.ResolveDependencies(pkgName)
	if err != nil {
		ui.Error("Bağımlılık analizi başarısız: %v", err)
		return err
	}
	
	if err := i.showInstallPlan(plan); err != nil {
		return err
	}
	
	if !i.cfg.NoConfirm {
		if !i.askConfirmation("Kuruluma devam edilsin mi?") {
			ui.Warning("Kurulum iptal edildi")
			return nil
		}
	}
	
	if len(plan.RepoPackages) > 0 {
		ui.Header("Resmi Repo Paketleri")
		ui.Info("Resmi repo paketleri kuruluyor...")
		if err := i.pacmanClient.Install(plan.RepoPackages, i.cfg.NoConfirm); err != nil {
			ui.Error("Repo paketleri kurulumu başarısız: %v", err)
			return err
		}
		ui.Success("Repo paketleri kuruldu")
	}
	
	if len(plan.AURPackages) > 0 {
		ui.Header("AUR Paketleri")
		for idx, aurPkg := range plan.AURPackages {
			ui.Info("AUR paketi kuruluyor (%d/%d): %s", idx+1, len(plan.AURPackages), aurPkg)
			if err := i.installAURPackage(aurPkg); err != nil {
				ui.Error("AUR paketi kurulumu başarısız (%s): %v", aurPkg, err)
				return err
			}
		}
	}
	
	ui.Separator()
	ui.Success("Tüm paketler başarıyla kuruldu!")
	return nil
}

func (i *Installer) showInstallPlan(plan *deps.InstallPlan) error {
	ui.InstallPlanBox(plan.RepoPackages, plan.AURPackages)
	return nil
}

func (i *Installer) askConfirmation(message string) bool {
	fmt.Printf("%s [E/h] (varsayılan: E): ", message)
	
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}
	
	response = strings.ToLower(strings.TrimSpace(response))
	return response == "" || response == "e" || response == "evet" || response == "y" || response == "yes"
}

func (i *Installer) installAURPackage(pkgName string) error {
	pkg, err := i.aurClient.Info(pkgName)
	if err != nil {
		return fmt.Errorf("paket bilgisi alınamadı: %w", err)
	}
	
	buildDir := filepath.Join(i.cfg.CacheDir, pkg.Name)
	if err := os.MkdirAll(buildDir, 0755); err != nil {
		return err
	}
	
	ui.Info("Kaynak kodu indiriliyor...")
	if err := i.cloneRepo(pkg.Name, buildDir); err != nil {
		return err
	}
	
	if !i.cfg.SkipReview {
		if !i.cfg.NoConfirm {
			if i.askConfirmation(fmt.Sprintf("%s PKGBUILD dosyasını incelemek istiyor musunuz?", pkg.Name)) {
				i.showPKGBUILD(buildDir)
			}
		}
	}
	
	ui.Info("Paket derleniyor ve kuruluyor...")
	if err := i.buildAndInstallPackage(buildDir); err != nil {
		return err
	}
	
	if i.cfg.CleanAfter {
		os.RemoveAll(buildDir)
	}
	
	ui.Success(fmt.Sprintf("%s kuruldu", pkg.Name))
	return nil
}

func (i *Installer) cloneRepo(pkgName, dest string) error {
	url := fmt.Sprintf("https://aur.archlinux.org/%s.git", pkgName)
	
	if _, err := os.Stat(dest); err == nil {
		os.RemoveAll(dest)
	}
	
	cmd := exec.Command(i.cfg.GitPath, "clone", url, dest)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func (i *Installer) showPKGBUILD(buildDir string) {
	pkgbuildPath := filepath.Join(buildDir, "PKGBUILD")
	
	cmd := exec.Command("less", pkgbuildPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func (i *Installer) buildAndInstallPackage(buildDir string) error {
	args := []string{"-si"}
	if i.cfg.NoConfirm {
		args = append(args, "--noconfirm")
	}
	
	cmd := exec.Command(i.cfg.MakepkgPath, args...)
	cmd.Dir = buildDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = append(os.Environ(), "MAKEFLAGS="+i.cfg.MakeFlags)
	
	return cmd.Run()
}
