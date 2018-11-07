package bruteforce

import (
	"regexp/syntax"

	"github.com/alixaxel/genex"
)

func Words(input, charset *syntax.Regexp) <-chan string {
	ch := make(chan string)
	go func() {
		genex.Generate(input, charset, 3, func(output string) {
			ch <- output
		})
		close(ch)
	}()
	return ch
}
