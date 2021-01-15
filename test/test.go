package main

import (
	"gopkg.in/yaml.v2"
	"fmt"
)

func main() {
	yml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

fmt.Println(yml)
	var redirectors []Redirector
	yaml.Unmarshal([]byte(yml), &redirectors)

	structToMap := map[string]string{}

	for _, redirector := range redirectors {
		fmt.Println(redirector)
		structToMap[redirector.Path] = redirector.URL
	}
	fmt.Println(structToMap)
}

type Redirector struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}