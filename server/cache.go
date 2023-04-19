package main

import (
	"net/http"
)

func main() {
	go func() {
		http.Handle("/", http.StripPrefix("/image", http.FileServer(http.Dir("./image/"))))
		http.ListenAndServe(":8080", nil)
	}()
	select {}
}
