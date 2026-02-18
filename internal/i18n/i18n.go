package i18n

type Translator struct {
	lang     string
	messages map[string]string
}

func New(lang string) *Translator {
	if lang != "tr" && lang != "en" {
		lang = "tr"
	}
	
	return &Translator{
		lang:     lang,
		messages: getMessages(lang),
	}
}

func (t *Translator) Get(key string, args ...interface{}) string {
	if msg, ok := t.messages[key]; ok {
		return msg
	}
	return key
}

func getMessages(lang string) map[string]string {
	if lang == "en" {
		return map[string]string{
			"searching":        "[lt] AUR'da aranıyor: %s...",
			"search_failed":    "Arama başarısız: %v",
			"no_results":       "Paket bulunamadı",
			"install_hint":     "Kurulum için: lt kur <isim>",
			"installing":       "[lt] '%s' kuruluyor...",
			"install_failed":   "Kurulum başarısız: %v",
			"install_complete": "[+] Kurulum tamamlandı: %s",
			"checking_updates": "[lt] AUR paketleri kontrol ediliyor...",
			"update_failed":    "Güncelleme başarısız: %v",
		}
	}
	
	return map[string]string{
		"searching":        "[lt] AUR'da aranıyor: %s...",
		"search_failed":    "Arama başarısız: %v",
		"no_results":       "Paket bulunamadı",
		"install_hint":     "Kurulum için: lt kur <isim>",
		"installing":       "[lt] '%s' kuruluyor...",
		"install_failed":   "Kurulum başarısız: %v",
		"install_complete": "[+] Kurulum tamamlandı: %s",
		"checking_updates": "[lt] AUR paketleri kontrol ediliyor...",
		"update_failed":    "Güncelleme başarısız: %v",
	}
}
