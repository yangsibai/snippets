/*
 * Rendering Plain Text
 * From: <http://www.alexedwards.net/blog/golang-response-snippets#plain>
 */

package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
