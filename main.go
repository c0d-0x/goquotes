package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/* {"id":507,"quote":"Let silence take you to the core of life.","author":"Rumi"}% */

type Quote_t struct {
	Id     int64  `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

func Quotes_gen(api string) *Quote_t {
	resp, err := http.Get(api)
	if err != nil {
		fmt.Println("Error: Failed to fetch quote.", err)
		return nil
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: Failed to read resp: ", err)
		return nil
	}

	var quote Quote_t
	errj := json.Unmarshal([]byte(bodyBytes), &quote)
	if errj != nil {
		fmt.Println("Error: Failed to Unmarshal resp: ", errj)
		return nil
	}

	return &quote
}

func Fmt_quote(quote *Quote_t) string {
	str := "%s\n\t\t\t%s"
	str = fmt.Sprintf(str, quote.Quote, quote.Author)
	return str
}

func main() {
	api := "https://dummyjson.com/quotes/random"

	quote := Quotes_gen(api)

	quote_fmt_str := Fmt_quote(quote)
	fmt.Println(quote_fmt_str)
}
