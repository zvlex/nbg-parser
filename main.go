package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Channel struct {
	XMLName xml.Name `xml:"rss"`
	Title   string   `xml:"channel>title"`
	Link    string   `xml:"channel>link"`
	Item    Item     `xml:"channel>item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	Guid        string `xml:"guid"`
}

type Table struct {
	Rows []Row `xml:"tr"`
}

type Row struct {
	Cols []string `xml:"td"`
}

type Currency struct {
	Name        string
	Description string
	Rate        string
	Img         string
	Index       string
}

const (
	NBG_URL string = "http://www.nbg.ge/rss.php"
)

func main() {
	data := parseData()

	channel := Channel{}
	unmarshalXml([]byte(data), &channel)

	table := Table{}
	dec := descriptionDecoder([]byte(channel.Item.Description))

	if err := dec.Decode(&table); err != nil {
		log.Fatalln(err)
	}

	currency := currencyWrapper(&table)

	for _, value := range currency {
		fmt.Println(value.Name, value.Description, value.Rate, value.Img, value.Index)
	}
}

func currencyWrapper(table *Table) []Currency {
	currency := []Currency{}

	for _, value := range table.Rows {
		cols := value.Cols

		new_currency := Currency{cols[0], cols[1], cols[2], cols[3], cols[4]}
		currency = append(currency, new_currency)
	}

	return currency
}

func parseData() []byte {
	response, err := http.Get(NBG_URL)

	defer response.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}

	return readResponce(response.Body)
}

func descriptionDecoder(data []byte) *xml.Decoder {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.Strict = false
	decoder.AutoClose = xml.HTMLAutoClose
	decoder.Entity = xml.HTMLEntity

	return decoder
}

func readResponce(response_body io.Reader) []byte {
	body, err := ioutil.ReadAll(response_body)

	if err != nil {
		log.Fatalln(err)
	}

	return body
}

func unmarshalXml(data []byte, value interface{}) {
	if err := xml.Unmarshal(data, value); err != nil {
		log.Fatalln(err)
	}
}
