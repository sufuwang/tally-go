package server

import (
	"fmt"
	"net/http"
)

var domain string = "localhost"

type StartParameter struct {
	Domain    string
	HttpPort  int
	HttpsPort int
}

func Start(P StartParameter) {
	if P.Domain != "" {
		domain = P.Domain
	}
	go startHttps(P.HttpsPort)
	startHTTP(P.HttpPort)
}

func startHttps(port int) {
	url := fmt.Sprintf("%s:%d", domain, port)
	err := http.ListenAndServeTLS(url, "src/public/credential/cert.pem", "src/public/credential/key.pem", nil)
	if err != nil {
		fmt.Println("HTTPS: ", err)
	}
}

func startHTTP(port int) {
	url := fmt.Sprintf("%s:%d", domain, port)
	err := http.ListenAndServe(url, nil)
	if err != nil {
		fmt.Println("HTTP: ", err)
	}
}
