package main

import "testing"

func TestWords(t *testing.T) {

	for template, v := range map[string]struct{ count int }{
		"sample": {1},
		"sampl#": {10},
	} {
		var combinations []string
		for word := range words(template) {
			combinations = append(combinations, word)
		}
		if len(combinations) != v.count {
			t.Fatalf("combinations number: %d, needed %d", len(combinations), v.count)
		}
	}

}
