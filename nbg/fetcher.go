package nbg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const NBG_URL = "https://nbg.gov.ge/gw/api/ct/monetarypolicy/currencies/ka/json"

func fetchCurrencyRates(currentDate string) *Exchange {
	resp, err := http.Get(NBG_URL + "?date=" + currentDate)

	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}

	var exchange []*Exchange

	if err = json.Unmarshal(body, &exchange); err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	return exchange[0]

}

func FetchRates(params *RequestParams) CurrencyData {
	currentDate := params.date
	code := params.code

	cache := readCache(currentDate)

	if cache != nil {
		return cache.FilterByCode(code)
	}

	exchangeData := fetchCurrencyRates(currentDate)

	cache = storeCache(currentDate, exchangeData.Currencies)

	return cache.FilterByCode(code)
}
