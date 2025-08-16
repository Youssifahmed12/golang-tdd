package main

import "strings"

const (
	englishprfx = "Hello "
	spanishprfx = "Hola "
	frenchPrfx  = "Bonjour "
)

func main() {

}

func hello(name, lang string) string {
	return getLangPrefix(lang) + name + "!"
}

func getLangPrefix(lang string) string {
	lang = strings.ToLower(lang)
	switch lang {
	case "english":
		return englishprfx
	case "spanish":
		return spanishprfx
	case "french":
		return frenchPrfx
	default:
		return englishprfx
	}

}
