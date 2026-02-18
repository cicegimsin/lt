# lt - Sade AUR Paket Yöneticisi

<div align="center">

⚡ Hızlı, sade ve tamamen Türkçe AUR paket yöneticisi

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

</div>

## 📋 Özellikler

- ⚡ **Hızlı**: Go ile yazılmış, paralel işlemler
- 🇹🇷 **Türkçe**: Tamamen Türkçe arayüz
- 🎨 **Sade**: Renkli ve anlaşılır CLI çıktısı
- 📦 **Akıllı**: Otomatik bağımlılık çözümleme
- 🔄 **Güvenli**: Onay mekanizması ile güncelleme ve kaldırma
- 🛡️ **Kontrollü**: PKGBUILD güvenlik kontrolü

## 🚀 Kurulum

### Gereksinimler

- Go 1.21 veya üzeri
- git
- base-devel (make, gcc, vb.)
- pacman (Arch Linux)

### Hızlı Kurulum (Önerilen)

```bash
go install github.com/cicegimsin/lt@latest
```

Kurulum sonrası `lt` komutunu kullanabilirsiniz. Eğer komut bulunamazsa, `~/go/bin` dizinini PATH'e ekleyin:

```bash
echo 'export PATH=$PATH:~/go/bin' >> ~/.bashrc
source ~/.bashrc
```

### Kaynak Koddan Kurulum

```bash
git clone https://github.com/cicegimsin/lt.git
cd lt
make build
sudo make install
```

## 📖 Kullanım

### Paket Arama

```bash
lt ara neovim
```

Çıktı:
```
[lt] AUR'da aranıyor: neovim...

neovim 0.9.5-1 ★ 1250
   Vim-fork focused on extensibility and usability
   Güncelleme: 2024-01-15

Kurulum için: lt kur <isim>
```

### Paket Kurma

```bash
lt kur yazi
```

### Paket Güncelleme

```bash
lt güncelle
```

Güncellenebilir paketleri listeler ve onay ister:
```
[+] 3 paket güncellenebilir:
  - yazi (0.2.4-1 -> 0.2.5-1)
  - zoxide (0.9.2-1 -> 0.9.4-1)

Güncellemek istiyor musunuz? [E/h]:
```

### Paket Kaldırma

```bash
lt kaldır paket-adi
```

Onay ister:
```
'paket-adi' paketini kaldırmak istediğinize emin misiniz? [E/h]:
```

### Diğer Komutlar

```bash
lt liste          # Kurulu AUR paketlerini listele
lt bilgi neovim   # Paket detaylarını göster
lt temizle        # Önbellek temizle
lt indir yazi     # Sadece kaynak kodunu indir
lt yardim         # Yardım menüsü
```

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

## 🎯 Örnekler

### Birden fazla paket kurma

```bash
lt kur yazi
lt kur zoxide
lt kur lf
```

### Paket bilgisi görüntüleme

```bash
lt bilgi neovim
```

Çıktı:
```
neovim 0.9.5-1
Açıklama: Vim-fork focused on extensibility and usability
Adres: https://neovim.io
Oy Sayısı: 1250
Popülerlik: 15.42
Bakımcı: username
Son Güncelleme: 2024-01-15 10:30
```

### Kurulu AUR paketlerini listeleme

```bash
lt liste
```

## 🔧 Sorun Giderme

### "lt: command not found" hatası

Go binary dizinini PATH'e ekleyin:

```bash
export PATH=$PATH:~/go/bin
```

Kalıcı yapmak için `.bashrc` veya `.zshrc` dosyanıza ekleyin.

### Derleme hatası

Bağımlılıkları güncelleyin:

```bash
go mod tidy
make build
```

### Önbellek sorunları

Önbelleği temizleyin:

```bash
lt temizle
```

## 🤝 Katkıda Bulunma

Katkılarınızı bekliyoruz! Lütfen:

1. Fork yapın
2. Feature branch oluşturun (`git checkout -b yeni-ozellik`)
3. Değişikliklerinizi commit edin (`git commit -am 'Yeni özellik eklendi'`)
4. Branch'inizi push edin (`git push origin yeni-ozellik`)
5. Pull Request oluşturun

## 📝 Lisans

Bu proje MIT lisansı altında lisanslanmıştır. Detaylar için [LICENSE](LICENSE) dosyasına bakın.

## 🙏 Teşekkürler

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Color](https://github.com/fatih/color) - Renkli terminal çıktısı
- AUR topluluğu

## 📧 İletişim

Sorularınız veya önerileriniz için issue açabilirsiniz.

---

<div align="center">
Made with ❤️ for Arch Linux users
</div>
