package main

import (
	"fmt"
	"net/http"
)
var count int
func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/requests", Requests)
	http.HandleFunc("/queries", Queries)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")

}

func Requests(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
	count = count + 1


}

func Queries(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, count)
}