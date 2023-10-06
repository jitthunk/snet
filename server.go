package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func GetFiles(path string) []string {
	files, err := os.ReadDir(path)

	var result []string
	if err == nil {

		for _, f := range files {
			result = append(result, f.Name())
		}
	}

	return result
}

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
	fmt.Printf("Server starting on %s 8080", GetOutboundIP())
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
