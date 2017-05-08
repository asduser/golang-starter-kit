package main

import (
	"fmt"
	"regexp"
)

var imgPattern = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)

func parseImages(html string) []string {
	images := imgPattern.FindAllStringSubmatch(html, -1)
	result := make([]string, len(images))
	for i := range result {
		result[i] = images[i][1]
	}
	return result
}

func main() {
	fmt.Printf("%q", parseImages(html))
}

const html = `
<img src="1.png"><x><z?>
<img czx zcxz src='2.png'><x><z?>
`