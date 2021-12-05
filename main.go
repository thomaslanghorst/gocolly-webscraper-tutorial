package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type EtfInfo struct {
	Title              string
	Replication        string
	Earnings           string
	TotalExpenseRatio  string
	TrackingDifference string
	FundSize           string
}

func main() {

	isins := []string{"IE00B1XNHC34", "IE00B4L5Y983", "LU1838002480"}

	etfInfo := EtfInfo{}
	etfInfos := make([]EtfInfo, 0, 1)

	c := colly.NewCollector(colly.AllowedDomains("www.trackingdifferences.com", "trackingdifferences.com"))

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US;q=0.9")
		fmt.Printf("Visiting %s\n", r.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error while scraping: %s\n", e.Error())
	})

	c.OnHTML("h1.page-title", func(h *colly.HTMLElement) {
		etfInfo.Title = h.Text
	})

	c.OnHTML("div.descfloat p.desc", func(h *colly.HTMLElement) {
		selection := h.DOM

		childNodes := selection.Children().Nodes
		if len(childNodes) == 3 {
			description := cleanDesc(selection.Find("span.desctitle").Text())
			value := selection.FindNodes(childNodes[2]).Text()

			switch description {
			case "Replication":
				etfInfo.Replication = value
				break
			case "TER":
				etfInfo.TotalExpenseRatio = value
				break
			case "TD":
				etfInfo.TrackingDifference = value
				break
			case "Earnings":
				etfInfo.Earnings = value
				break
			case "Fund size":
				etfInfo.FundSize = value
				break
			}
		}
	})

	c.OnScraped(func(r *colly.Response) {
		etfInfos = append(etfInfos, etfInfo)
		etfInfo = EtfInfo{}
	})

	for _, isin := range isins {
		c.Visit(scrapeUrl(isin))
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(etfInfos)

}

func cleanDesc(s string) string {
	return strings.TrimSpace(s)
}

func scrapeUrl(isin string) string {
	return "https://www.trackingdifferences.com/ETF/ISIN/" + isin
}
