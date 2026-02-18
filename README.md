# lt - Sade AUR Paket Yöneticisi

Hızlı, sade ve Türkçe destekli AUR paket yöneticisi.

## Özellikler

- ⚡ Hızlı arama ve kurulum (Go ile yazılmış)
- 🇹🇷 Türkçe arayüz
- 🎨 Renkli ve sade CLI çıktısı
- 📦 Otomatik bağımlılık çözümleme
- 🔄 Paralel güncelleme desteği
- 🛡️ PKGBUILD güvenlik kontrolü

## Kurulum

### Kaynak Koddan

```bash
git clone https://github.com/cicegimsin/lt-aur-helper.git
cd lt
make build
sudo make install
```

### Hızlı Kurulum

```bash
go install https://github.com/cicegimsin/lt-aur-helper@latest
```

## Kullanım

### Paket Arama
```bash
lt ara neovim
```

### Paket Kurma
```bash
lt kur yazi
```

### Güncelleme
```bash
lt güncelle
```

### Diğer Komutlar
```bash
lt liste          # Kurulu AUR paketlerini listele
lt bilgi <paket>  # Paket detayları
lt kaldır <paket> # Paket kaldır
lt temizle        # Önbellek temizle
lt indir <paket>  # Sadece kaynak indir
```

## Gereksinimler

- Go 1.21+
- git
- base-devel (make, gcc, vb.)
- pacman

## Lisans

MIT/home/click/Masaüstü/lt/README.md
