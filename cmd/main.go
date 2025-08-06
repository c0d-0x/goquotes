package main

import (
	"fmt"

	"github.com/c0d-0x/goquotes/quote_gen"
)

func main() {
	api := "https://dummyjson.com/quotes/random"

	quote := quote_gen.Quotes_gen(api)

	fmt.Println(quote_gen.Fmt_qoute)
}
