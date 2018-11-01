package main

import "fmt"

var smaz []byte
var lgaz []byte
var dgts []byte

func init() {
	var src []byte
	for i := 0; i < 123; i++ {
		src = append(src, byte(i))
	}
	dgts = src[48:58]
	smaz = src[97:123]
	lgaz = src[65:91]
}

func words() <-chan string {
	ch := make(chan string)
	go func() {
		for sample := range generate(dgts, 1) {
			ch <- fmt.Sprintf("passwor%s", sample)
		}
		close(ch)
	}()
	return ch
}
