package hello

import "rsc.io/quote"

func Hello() string {
	return "Hello, Gopher!"
}

func Hello_quote() string {
	return quote.Hello()
}
