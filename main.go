package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Quote_t struct {
	Id     int64  `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

func quotes_gen(api string) *Quote_t {
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

func fmt_quote(quote *Quote_t) string {
	str := "\"%s\"\n >>%s<<"
	str = fmt.Sprintf(str, quote.Quote, quote.Author)
	return str
}

func draw_ascii_art() {
	ascii_art := "\n /\\_/\\\n" +
		"( o.o )\n" +
		" > - < "

	fmt.Println(ascii_art)
}

func main() {
	api := "https://dummyjson.com/quotes/random"

	quote := quotes_gen(api)

	quote_fmt_str := fmt_quote(quote)
	draw_ascii_art()
	fmt.Println(quote_fmt_str)
}
