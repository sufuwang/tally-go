package main

import (
	"net/http"
	"tally-go/src/server"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	server.Start(server.StartParameter{
		Domain:    "localhost",
		HttpPort:  8000,
		HttpsPort: 1443,
	})
}
