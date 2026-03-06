# lt - Evrensel Paket Yöneticisi

<div align="center">

🌍 Tüm işletim sistemleri için birleşik paket yöneticisi

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

[Kurulum](#kurulum) • [Kullanım](#kullanım) • [Desteklenen Sistemler](#desteklenen-sistemler)

</div>

---

## Özellikler

- 🌍 **Evrensel Destek** - Linux, macOS, Windows
- 🔍 **Akıllı Tespit** - Otomatik işletim sistemi ve paket yöneticisi tespiti
- 🎨 **Renkli CLI** - Terminal genişliğine göre otomatik ayarlama
- 📦 **Çoklu Kaynak** - Resmi repolar + AUR + Homebrew + Chocolatey
- 🇹🇷 **Türkçe Arayüz** - Komutlar ve çıktılar tamamen Türkçe
- ⚡ **Hızlı** - Go ile yazılmış, paralel işlemler

---

## Desteklenen Sistemler

### Linux Dağıtımları
- **Arch Linux** - pacman + AUR desteği
- **Ubuntu/Debian** - apt paket yöneticisi
- **Fedora/CentOS/RHEL** - dnf/yum paket yöneticisi
- **openSUSE** - zypper paket yöneticisi
- **Alpine Linux** - apk paket yöneticisi

### macOS
- **Homebrew** - brew paket yöneticisi
- **MacPorts** - port paket yöneticisi

### Windows
- **Chocolatey** - choco paket yöneticisi
- **Scoop** - scoop paket yöneticisi

---

## Kurulum

```bash
# Tüm sistemler için
go install github.com/cicegimsin/lt@latest

# PATH ayarı (gerekirse)
echo 'export PATH=$PATH:~/go/bin' >> ~/.bashrc
source ~/.bashrc
```

---

## Kullanım

### Temel Komutlar

```bash
lt ara firefox          # Paket ara (tüm kaynaklardan)
lt kur firefox          # Paket kur
lt güncelle            # Sistem güncelle
lt kaldır firefox      # Paket kaldır
lt sistem              # Sistem bilgileri
```

### Örnek Çıktılar

**Sistem Tespiti:**
```
◆ SİSTEM BİLGİLERİ
═══════════════════════════════════════════════════════

▶ İŞLETİM SİSTEMİ
  İşletim Sistemi: Ubuntu Linux
  Dağıtım: ubuntu
  Sürüm: 22.04
  Mimari: linux

▶ PAKET YÖNETİCİSİ
  Paket Yöneticisi: apt
  Sudo Gerekli: true
  AUR Desteği: false
```

**Evrensel Arama:**
```
◆ PAKET ARAMA
═══════════════════════════════════════════════════════
→ Sistem: Ubuntu Linux 22.04 (apt)
→ Paket aranıyor: firefox

EXTRA Resmi Paketler (1 paket)
──────────────────────────────────────────────────────
  • firefox 108.0.1+build1-0ubuntu1
    Mozilla Firefox web browser

AUR (2 paket)
──────────────────────────────────────────────────────
  • firefox-beta 109.0b9-1 ★ 45
    Standalone web browser from mozilla.org - Beta
```

---

## Yapılandırma

Dosya: `~/.config/lt/config.toml`

```toml
language = "tr"
no_confirm = false

# Otomatik tespit edilen değerler
[system]
os_type = "linux"
distribution = "ubuntu"
package_manager = "apt"
sudo_path = "/usr/bin/sudo"
```

---

## Gelişmiş Özellikler

### Çoklu Paket Yöneticisi Desteği
- Sistem otomatik olarak mevcut paket yöneticilerini tespit eder
- Birden fazla paket yöneticisi varsa en uygun olanı seçer
- AUR desteği Arch tabanlı sistemlerde otomatik aktif olur

### Akıllı Arama
- Resmi repolardan ve alternatif kaynaklardan arama
- Sonuçları kaynaklarına göre kategorize etme
- Terminal genişliğine göre otomatik formatlar

### Platform Özel Optimizasyonlar
- **Linux**: sudo yetkilendirmesi, paket bağımlılıkları
- **macOS**: Homebrew/MacPorts entegrasyonu
- **Windows**: PowerShell/CMD uyumluluğu

---

## Lisans

MIT License - detaylar için [LICENSE](LICENSE) dosyasına bakın.

---

<div align="center">

**Tüm işletim sistemleri için ❤️ ile yapıldı**

</div>