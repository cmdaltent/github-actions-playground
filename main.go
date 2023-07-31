// Simply the main package. This is the entry point of the program.
package main

import (
	"net/http"
)

func main() {

	if err := http.ListenAndServe(":8080", http.HandlerFunc(handler)); err != nil {
		panic(err)
	}

}
