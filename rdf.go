package main

import (
	"fmt"

	"github.com/Callidon/joseki/rdf"
)

func Triple() rdf.Triple {
	subject := rdf.NewURI("http://example.org/book/book1")
	predicate := rdf.NewURI("http://purl.org/dc/terms/title")
	object := rdf.NewLiteral("Harry Potter and the Order of the Phoenix")
	triple := rdf.NewTriple(subject, predicate, object)
	return triple
}

func main() {
	fmt.Println(Triple())
	// Output : {<http://example.org/book/book1> <http://purl.org/dc/terms/title> "Harry Potter and the Order of the Phoenix"}
}
