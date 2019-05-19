package main

import (
	"fmt"
	"testing"
)

func TestTriple(t *testing.T) {
	got := fmt.Println(Triple())
	want := "{<http://example.org/book/book1> <http://purl.org/dc/terms/title> \"Harry Potter and the Order of the Phoenix\"}"
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
