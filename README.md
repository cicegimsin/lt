# lt - Türkçe AUR Yardımcısı

<div align="center">

⚡ Hızlı ve sade Türkçe AUR paket yöneticisi

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

[Kurulum](#kurulum) • [Kullanım](#kullanım) • [Özellikler](#özellikler)

</div>

---

## Özellikler

- 🇹🇷 **Türkçe arayüz** - Komutlar ve çıktılar tamamen Türkçe
- ⚡ **Hızlı** - Go ile yazılmış, paralel işlemler
- � **Renkli çıktı** - Terminall'de güzel görünüm
- 📦 **Akıllı bağımlılık çözümü** - Otomatik dependency management
- 🏛️ **Hibrit arama** - Hem resmi repo hem AUR
- 🔧 **Pacman entegrasyonu** - Sistem ile tam uyumlu

---
go install github.com/cicegimsin/lt@latest
## Kurulum

```bash
# Gereksinimler
sudo pacman -S go git base-devel

# lt kurulumu
go install github.com/cicegimsin/lt@latest

# PATH ayarı
echo 'export PATH=$PATH:~/go/bin' >> ~/.bashrc
source ~/.bashrc
```

---

## Kullanım

### Temel komutlar

```bash
lt ara neovim          # Paket ara
lt kur yazi            # Paket kur
lt güncelle            # Sistem + AUR güncelle
lt kaldır paket-adi    # Paket kaldır
lt liste               # Kurulu paketleri listele
lt bilgi neovim        # Paket bilgisi
lt temizle             # Önbellek temizle
```

### Seçenekler

```bash
lt kur --noconfirm paket    # Onaysız kurulum
```
go install github.com/cicegimsin/lt@latest
### Örnek çıktılar

**Hibrit arama:**
```
$ lt ara firefox

◆ Resmi Repolar (1 paket)
extra/firefox 121.0-1
    Free and open-source web browser from Mozilla

◆ AUR (3 paket)
aur/firefox-beta 122.0b9-1 ★ 45
    Standalone web browser from mozilla.org - Beta
    Son güncelleme: 15.12.2023

─────────────────────────────────────────
→ Toplam 4 paket bulundu
```

**Akıllı kurulum:**
```
$ lt kur yazi

→ Bağımlılıklar analiz ediliyor...

◆ Kurulum Planı
  repo Resmi repo paketleri:
    • gcc
    • make

  aur AUR paketleri:
    • yazi

→ Toplam 3 paket kurulacak
Kuruluma devam edilsin mi? [E/h] (varsayılan: E): 

◆ Resmi repo paketleri kuruluyor
✓ Repo paketleri kuruldu

◆ AUR paketi (1/1): yazi
→ Kaynak kodu indiriliyor...
→ Paket derleniyor ve kuruluyor...
✓ yazi kuruldu

─────────────────────────────────────────
✓ Tüm paketler başarıyla kuruldu!
```

---

## Yapılandırma

Dosya: `~/.config/lt/config.toml`

```toml
language = "tr"
makeflags = "-j$(nproc)"
cache_dir = "~/.cache/lt"

# Sistem yolları
pacman_path = "/usr/bin/pacman"
sudo_path = "/usr/bin/sudo"
git_path = "/usr/bin/git"
makepkg_path = "/usr/bin/makepkg"

# Davranış
skip_review = false
clean_after = true
no_confirm = false
```

---

## Lisans

MIT License - detaylar için [LICENSE](LICENSE) dosyasına bakın.

---

## Katkı

Pull request'ler ve issue'lar hoş geldiniz.

---

<div align="center">

Arch Linux kullanıcıları için Click tarafından yapıldı

</div>
