# Kurulum Rehberi

## Hızlı Kurulum (Önerilen)

### 1. Go Kurulumu

Arch Linux için:
```bash
sudo pacman -S go git base-devel
```

### 2. lt Kurulumu

```bash
go install github.com/cicegimsin/lt@latest
```

### 3. PATH Ayarı

Eğer `lt: command not found` hatası alırsanız:

**Bash için:**
```bash
echo 'export PATH=$PATH:~/go/bin' >> ~/.bashrc
source ~/.bashrc
```

**Zsh için:**
```bash
echo 'export PATH=$PATH:~/go/bin' >> ~/.zshrc
source ~/.zshrc
```

**Fish için:**
```bash
fish_add_path ~/go/bin
```

### 4. Test

```bash
lt --version
lt ara yay
```

---

## Manuel Kurulum (Kaynak Koddan)

### 1. Gereksinimler

```bash
sudo pacman -S go git base-devel
```

### 2. Projeyi Klonla

```bash
git clone https://github.com/cicegimsin/lt.git
cd lt
```

### 3. Derle ve Kur

```bash
make build
sudo make install
```

### 4. Test

```bash
lt --version
```

---

## Güncelleme

### go install ile kurduysanız:
```bash
go install github.com/cicegimsin/lt@latest
```

### Manuel kurduysanız:
```bash
cd lt
git pull
make build
sudo make install
```

---

## Kaldırma

### go install ile kurduysanız:
```bash
rm ~/go/bin/lt
```

### Manuel kurduysanız:
```bash
sudo rm /usr/local/bin/lt
```

---

## Sorun Giderme

### "lt: command not found"

PATH'e ekleyin:
```bash
export PATH=$PATH:~/go/bin
```

### "permission denied"

sudo ile çalıştırın:
```bash
sudo make install
```

### Derleme hatası

Bağımlılıkları güncelleyin:
```bash
go mod tidy
make build
```
