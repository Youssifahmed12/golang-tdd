package main

func main() {

}

func hello(name, lang string) string {
	if lang == "Spanish" {
		return "Hola " + name + "!"
	}
	if lang == "French" {
		return "Bonjour " + name + "!"
	}
	return "Hello " + name + "!"
}
