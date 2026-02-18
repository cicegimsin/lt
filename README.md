# lt - Sade AUR Paket Yöneticisi

<div align="center">

⚡ Hızlı, sade ve tamamen Türkçe AUR paket yöneticisi

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

[Kurulum](#-kurulum) • [Kullanım](#-kullanım) • [Özellikler](#-özellikler) • [Katkıda Bulun](CONTRIBUTING.md)

</div>

---

## 📋 Özellikler

- ⚡ **Hızlı**: Go ile yazılmış, paralel işlemler
- 🇹🇷 **Türkçe**: Tamamen Türkçe arayüz
- 🎨 **Sade**: Renkli ve anlaşılır CLI çıktısı
- 📦 **Akıllı**: Otomatik bağımlılık çözümleme
- 🔄 **Güvenli**: Onay mekanizması ile güncelleme ve kaldırma
- 🛡️ **Kontrollü**: PKGBUILD güvenlik kontrolü

---

## 🚀 Kurulum

### Hızlı Kurulum (Önerilen)

```bash
# Go kurulumu (eğer yoksa)
sudo pacman -S go git base-devel

# lt kurulumu
go install github.com/cicegimsin/lt@latest

# PATH'e ekle (gerekirse)
echo 'export PATH=$PATH:~/go/bin' >> ~/.bashrc
source ~/.bashrc
```

### Detaylı Kurulum

Detaylı kurulum talimatları için [INSTALL.md](INSTALL.md) dosyasına bakın.

---

## 📖 Kullanım

### Temel Komutlar

```bash
# Paket ara
lt ara neovim

# Paket kur
lt kur yazi

# Paketleri güncelle
lt güncelle

# Paket kaldır
lt kaldır paket-adi

# Kurulu paketleri listele
lt liste

# Paket bilgisi
lt bilgi neovim

# Önbellek temizle
lt temizle

# Yardım
lt yardim
```

### Örnek Kullanım

**Paket Arama:**
```bash
$ lt ara yay

[lt] AUR'da aranıyor: yay...

yay 12.3.5-1 ★ 2450
   Yet another yogurt. Pacman wrapper and AUR helper written in go.
   Güncelleme: 2024-01-20

Kurulum için: lt kur <isim>
```

**Paket Güncelleme:**
```bash
$ lt güncelle

[lt] AUR paketleri kontrol ediliyor...

[+] 2 paket güncellenebilir:
  - yazi (0.2.4-1 -> 0.2.5-1)
  - zoxide (0.9.2-1 -> 0.9.4-1)

Güncellemek istiyor musunuz? [E/h] (varsayılan: E): 

[1/2] yazi güncelleniyor...
[+] yazi güncellendi (0.2.5-1)
[2/2] zoxide güncelleniyor...
[+] zoxide güncellendi (0.9.4-1)

[+] Güncelleme tamamlandı!
```

---

## ⚙️ Yapılandırma

Yapılandırma dosyası otomatik oluşturulur: `~/.config/lt/config.toml`

```toml
language = "tr"
makeflags = "-j$(nproc)"
parallel_downloads = 5
color_scheme = "default"
cache_dir = "~/.cache/lt"
log_dir = "~/.local/share/lt/logs"
```

---

## 🔧 Sorun Giderme

### "lt: command not found" hatası

```bash
export PATH=$PATH:~/go/bin
```

Kalıcı yapmak için `.bashrc` veya `.zshrc` dosyanıza ekleyin.

### Derleme hatası

```bash
go mod tidy
make build
```

### Daha fazla yardım

[INSTALL.md](INSTALL.md) dosyasına bakın veya [issue açın](https://github.com/cicegimsin/lt/issues).

---

## 🤝 Katkıda Bulunma

Katkılarınızı bekliyoruz! [CONTRIBUTING.md](CONTRIBUTING.md) dosyasına göz atın.

---

## 📝 Lisans

Bu proje MIT lisansı altında lisanslanmıştır. Detaylar için [LICENSE](LICENSE) dosyasına bakın.

---

## 🙏 Teşekkürler

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Color](https://github.com/fatih/color) - Renkli terminal çıktısı
- AUR topluluğu

---

<div align="center">

**[⬆ Başa Dön](#lt---sade-aur-paket-yöneticisi)**

Made with ❤️ for Arch Linux users

</div>

