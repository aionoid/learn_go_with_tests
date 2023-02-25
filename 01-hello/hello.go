// @Title
// @Description
// @Author
// @Update
package hello

import (
	"fmt"
)

const helloEnglishPrefix = "Hello, %s!"
const helloFrenchPrefix = "Bonjour, %s!"
const helloSpanishPrefix = "Holla, %s!"
const fr = "fr"
const sp = "sp"
const en = "en"

func Hello(name, language string) (msg string) {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf(getlangPrefix(language), name)
}

func getlangPrefix(lang string) (prefix string) {
	switch lang {
	case fr:
		prefix = helloFrenchPrefix
	case sp:
		prefix = helloSpanishPrefix
	default:
		prefix = helloEnglishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", "eng"))
}
