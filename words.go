package main

import (
	"regexp/syntax"

	"github.com/alixaxel/genex"
)

func words(resume string, input, charset *syntax.Regexp) <-chan string {
	ch, rawCh := make(chan string), make(chan string)
	go func() {
		genex.Generate(input, charset, 3, func(output string) {
			rawCh <- output
		})
		close(rawCh)
	}()
	go func() {
		var resumeReached bool
		for word := range rawCh {
			if resume != "" {
				if resumeReached {
					ch <- word
				} else if word == resume {
					resumeReached = true
					ch <- word
				}
			} else {
				ch <- word
			}
		}
		close(ch)
	}()
	return ch
}
