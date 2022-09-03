package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exeDir := filepath.Dir(exePath)
	fs := http.FileServer(http.Dir(exeDir + "/public"))
	http.Handle("/", fs)

	http.HandleFunc("/404", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, exeDir + "/public/404.html")
	})

	log.Print("Listening on :5000...")
	er := http.ListenAndServe(":5000", nil)
	if er != nil {
		log.Fatal(er)
	}
}
