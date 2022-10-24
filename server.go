package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	//path := r.URL.Path[1:]
	base := "."
	tmpl := "<a href=\"%s\">%s</a><br/>"
	files := GetFiles(base)
	var resp string
	for _, n := range files {
		resp += fmt.Sprintf(tmpl, n, n)
	}
	fmt.Fprintf(w, "<html>%s</html>", resp)
}
func main() {
	fmt.Print("Server starting")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
