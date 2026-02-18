# Katkıda Bulunma Rehberi

lt projesine katkıda bulunmak istediğiniz için teşekkürler! 🎉

## Geliştirme Ortamı Kurulumu

### Gereksinimler

- Go 1.21 veya üzeri
- git
- make
- Arch Linux veya Arch tabanlı dağıtım (test için)

### Projeyi Klonlama

```bash
git clone https://github.com/cicegimsin/lt.git
cd lt
```

### Bağımlılıkları Yükleme

```bash
go mod download
```

### Derleme

```bash
make build
```

### Test Etme

```bash
./lt ara yay
./lt bilgi neovim
```

## Kod Standartları

### Go Kod Stili

- `gofmt` kullanın
- Değişken isimleri Türkçe olabilir ama fonksiyon isimleri İngilizce olmalı
- Her fonksiyon için yorum satırı ekleyin

### Commit Mesajları

Türkçe commit mesajları kullanın:

```
✨ Yeni özellik: Paralel güncelleme desteği
🐛 Düzeltme: AUR API timeout sorunu
📝 Dokümantasyon: README güncellendi
♻️ Refactor: Bağımlılık çözümleme iyileştirildi
```

## Pull Request Süreci

1. Feature branch oluşturun: `git checkout -b yeni-ozellik`
2. Değişikliklerinizi yapın
3. Test edin: `make build && ./lt`
4. Commit edin: `git commit -am "Yeni özellik eklendi"`
5. Push edin: `git push origin yeni-ozellik`
6. Pull Request açın

## Yeni Özellik Ekleme

### Yeni Komut Ekleme

1. `cmd/` dizininde yeni dosya oluşturun
2. Cobra command yapısını kullanın
3. Türkçe açıklamalar ekleyin
4. `cmd/root.go` içinde init edin

Örnek:

```go
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/cicegimsin/lt/internal/ui"
)

var yeniCmd = &cobra.Command{
	Use:     "yeni [argüman]",
	Aliases: []string{"new"},
	Short:   "Yeni özellik açıklaması",
	Run: func(cmd *cobra.Command, args []string) {
		ui.Info("Yeni özellik çalışıyor")
	},
}

func init() {
	rootCmd.AddCommand(yeniCmd)
}
```

## Test Etme

### Manuel Test

```bash
make build
./lt ara test
./lt kur test-paketi
./lt güncelle
```

### Temizlik

```bash
make clean
```

## Sorular?

Issue açarak sorabilirsiniz!
