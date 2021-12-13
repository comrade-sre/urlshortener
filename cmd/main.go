package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	urlShortener := &http.Server {
		Addr:	":8080",
		Handler: mainHandler,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(urlShortener.ListenAndServe())

}
