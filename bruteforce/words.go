package bruteforce

import (
	"regexp/syntax"

	"github.com/alixaxel/genex"
)

func Words(resume string, input, charset *syntax.Regexp) <-chan string {
	ch, rawCh := make(chan string), make(chan string)
	go func() {
		genex.Generate(input, charset, 3, func(output string) {
			rawCh <- output
		})
		close(rawCh)
	}()
	go func() {
		var next bool
		for word := range rawCh {
			if !next && word == resume {
				next = true
			}
			if resume != "" && next {
				ch <- word
			} else {
				ch <- word
			}
		}
		close(ch)
	}()
	return ch
}
