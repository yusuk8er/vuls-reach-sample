package main

import (
	"fmt"

	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

func main() {
	tags, _, err := language.ParseAcceptLanguage("ja,en-US;q=0.8")
	fmt.Println(tags, err)

	out, _ := yaml.Marshal(map[string]string{"hello": "world"})
	fmt.Println(string(out))
}
