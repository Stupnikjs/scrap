package main

import (
	"github.com/Stupnikjs/scrap/lkdn"
)

var lkdnURL = "https://www.linkedin.com"

func main() {
	lkdn.Scrap(lkdnURL)
}
