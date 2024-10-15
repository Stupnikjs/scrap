package main

import (
	"github.com/Stupnikjs/scrap/lkdn"
)

var lkdnURL = "https://www.linkedin.com/search/results/all/?fetchDeterministicClustersOnly=true&heroEntityKey=urn%3Ali%3Afsd_profile%3AACoAABSTHFEBxHw5-DUKJeIdJYfuUeQKYhPjAxQ&keywords=maxime%20jumelle&origin=RICH_QUERY_SUGGESTION&position=0&searchId=9113c1d0-0c0e-4e9a-a8d4-86717f8f97ab&sid=%3B~o&spellCorrectionEnabled=false"

func main() {
	lkdn.Scrap(lkdnURL)
}
