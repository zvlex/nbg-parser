package main

import (
	"flag"
	"github.com/jedib0t/go-pretty/v6/table"
	"nbgcurr/nbg"
	"os"
	"time"
)

var currencyCode, publishDate string

func init() {
	now := time.Now().Format("2006-01-02")

	flag.StringVar(&currencyCode, "code", "", "provide currency code")
	flag.StringVar(&publishDate, "date", now, "provide publish date")
	flag.Parse()
}

func main() {
	params := nbg.NewRequestParams(currencyCode, publishDate)

	rates := nbg.FetchRates(params)

	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Code", "Name", "Rate", "Date"})

	for _, c := range rates {
		t.AppendRow([]interface{}{c.Code, c.Name, c.Rate, publishDate})
	}

	t.Render()
}
