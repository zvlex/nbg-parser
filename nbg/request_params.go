package nbg

var currencyCode, publishDate string

type RequestParams struct {
	code, date string
}

func NewRequestParams(code, date string) *RequestParams {
	return &RequestParams{code, date}
}
