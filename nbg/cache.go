package nbg

type CurrencyData map[string]*Currency
type Cache map[string]map[string]*Currency

var rateCache Cache = make(Cache)

func storeCache(currentDate string, currencies []*Currency) CurrencyData {
	rateCache[currentDate] = make(CurrencyData)

	for _, c := range currencies {
		rateCache[currentDate][c.Code] = c
	}

	return rateCache[currentDate]
}

func readCache(currentDate string) CurrencyData {
	if cache, ok := rateCache[currentDate]; ok {
		return cache
	}

	return nil
}

func (c *CurrencyData) FilterByCode(code string) CurrencyData {
	var filteredCache CurrencyData = make(CurrencyData)

	for k, v := range *c {
		if k == code {
			filteredCache[k] = v

			return filteredCache
		}
	}

	return *c
}
