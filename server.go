package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {

	upath := r.URL.Path[1:]
	fpath := upath
	tpath := upath
	if len(upath) == 0 {
		fpath = basePath
	} else {
		fpath = filepath.Join(basePath, fpath)
		tpath = upath + "/"
	}

	log.Printf("Request Path = %s, File Path = %s", upath, fpath)
	f, err := os.Stat(fpath)
	if err == nil && !f.IsDir() {
		w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(f.Name()))
		w.Header().Set("Content-Type", "application/octet-stream")
		http.ServeFile(w, r, fpath)
		return
	} else {
		log.Printf("ERR on path %v", err)
	}

	tmpl := "<a href=\"%s\">%s</a><br/>"
	files := GetFiles(fpath)
	var resp string
	for _, n := range files {
		resp += fmt.Sprintf(tmpl, tpath+n, n)
	}
	fmt.Fprintf(w, "<html>%s</html>", resp)
}

var basePath = "."

func main() {
	if len(os.Args) > 1 {
		basePath = os.Args[1]
	}
	fmt.Print("Server starting")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
