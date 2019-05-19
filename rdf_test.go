package main

import (
	"testing"

	"github.com/Callidon/joseki/rdf"
)

func TestTriple(t *testing.T) {
	got := Triple()
	want := rdf.NewTriple(rdf.NewURI("http://example.org/book/book1"), rdf.NewURI("http://purl.org/dc/terms/title"), rdf.NewLiteral("Harry Potter and the Order of the Phoenix"))
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
