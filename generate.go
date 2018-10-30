package main

import "strings"

func generate(charMap []byte, l int) <-chan string {
	res := make(chan string)
	go process(res, charMap, make([]byte, l), l, 0)
	return res
}

func process(res chan string, charMap []byte, b []byte, l int, j int) {
	for i := 0; i < len(charMap); i++ {
		b[j] = charMap[i]
		if j+1 < l {
			process(res, charMap, b, l, j+1)
		} else {
			res <- string(b)
		}
	}
	if string(b) == strings.Repeat(string(charMap[len(charMap)-1]), l) && j == 0 {
		close(res)
	}
}
