package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	//path := r.URL.Path[1:]
	base := "."

	files := GetFiles(base)
	var resp string
	for _, n := range files {
		resp += n + "<br/>"
	}
	fmt.Fprintf(w, "<html>%s</html>", resp)
}
func main() {
	fmt.Print("Server starting")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
