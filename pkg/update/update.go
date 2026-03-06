package update

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cicegimsin/lt/internal/aur"
	"github.com/cicegimsin/lt/internal/config"
	"github.com/cicegimsin/lt/internal/i18n"
	"github.com/cicegimsin/lt/internal/pacman"
	"github.com/cicegimsin/lt/internal/ui"
	"github.com/cicegimsin/lt/pkg/install"
)

type Updater struct {
	cfg          *config.Config
	tr           *i18n.Translator
	aurClient    *aur.Client
	pacmanClient *pacman.Client
}

type UpdateInfo struct {
	Name       string
	LocalVer   string
	RemoteVer  string
	IsAUR      bool
}

func New(cfg *config.Config, tr *i18n.Translator) *Updater {
	return &Updater{
		cfg:          cfg,
		tr:           tr,
		aurClient:    aur.NewClient(),
		pacmanClient: pacman.NewClient(cfg.PacmanPath, cfg.SudoPath),
	}
}

func (u *Updater) Update() error {
	ui.Banner("SISTEM VE AUR GÜNCELLEMESİ")
	
	ui.Header("Sistem Güncellemesi")
	ui.Info("Sistem paketleri güncelleniyor...")
	
	if err := u.pacmanClient.Update(u.cfg.NoConfirm); err != nil {
		ui.Warning("Sistem güncellemesi başarısız: %v", err)
	} else {
		ui.Success("Sistem paketleri güncellendi")
	}
	
	ui.Header("AUR Paket Kontrolü")
	ui.Info("AUR paketleri kontrol ediliyor...")
	
	aurPackages, err := u.pacmanClient.GetForeignPackages()
	if err != nil {
		ui.Error("Kurulu AUR paketleri alınamadı: %v", err)
		return err
	}
	
	if len(aurPackages) == 0 {
		ui.Box("BİLGİ", "Kurulu AUR paketi bulunamadı")
		return nil
	}
	
	var updates []UpdateInfo
	ui.Info("Sürüm kontrolü yapılıyor...")
	
	for i, pkg := range aurPackages {
		ui.Progress(i, len(aurPackages), pkg.Name)
		
		aurPkg, err := u.aurClient.Info(pkg.Name)
		if err != nil {
			continue
		}
		
		if pkg.Version != aurPkg.Version {
			updates = append(updates, UpdateInfo{
				Name:      pkg.Name,
				LocalVer:  pkg.Version,
				RemoteVer: aurPkg.Version,
				IsAUR:     true,
			})
		}
	}
	fmt.Println() // Progress satırını bitir
	
	if len(updates) == 0 {
		ui.Success("Tüm AUR paketleri güncel")
		return nil
	}
	
	var updateItems []string
	for _, upd := range updates {
		updateItems = append(updateItems, fmt.Sprintf("%s  %s → %s", 
			ui.Bold(upd.Name), 
			ui.Highlight(upd.LocalVer),
			ui.Highlight(upd.RemoteVer)))
	}
	
	ui.CategoryBox("Güncellenebilir Paketler", fmt.Sprintf("%d paket", len(updates)), updateItems)
	
	if !u.cfg.NoConfirm {
		if !u.askConfirmation("AUR paketlerini güncellemek istiyor musunuz?") {
			ui.Box("İPTAL", "AUR güncellemesi iptal edildi")
			return nil
		}
	}
	
	ui.Header("AUR Paket Güncellemesi")
	installer := install.New(u.cfg, u.tr)
	
	for i, upd := range updates {
		ui.Info("Güncelleniyor (%d/%d): %s", i+1, len(updates), upd.Name)
		
		if err := installer.Install(upd.Name); err != nil {
			ui.Error("%s güncellenemedi: %v", upd.Name, err)
			continue
		}
		
		ui.Success("%s güncellendi (%s)", upd.Name, upd.RemoteVer)
	}
	
	ui.Separator()
	ui.Success("Tüm güncellemeler tamamlandı!")
	
	return nil
}

func (u *Updater) askConfirmation(message string) bool {
	fmt.Printf("%s [E/h] (varsayılan: E): ", message)
	
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}
	
	response = strings.ToLower(strings.TrimSpace(response))
	return response == "" || response == "e" || response == "evet" || response == "y" || response == "yes"
}
