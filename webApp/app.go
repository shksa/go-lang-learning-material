package main

import (
	"log"
	"net/http"

	"github.com/shksa/learningGo/webApp/homepage"
	"github.com/shksa/learningGo/webApp/homepageC"
)

func main() {
	http.HandleFunc("/home", homepage.Handler)
	http.HandleFunc("/homeC", homepageC.Handler)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}
