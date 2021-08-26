package main

import (
	"tally-go/src/database"
	"tally-go/src/router"
	"tally-go/src/server"
)

func main() {
	database.LinkDataBase()
	router.RegisterRouter()
	server.Start(server.StartParameter{
		Domain:   "127.0.0.1",
		HttpPort: 3000,
		// HttpsPort: 1443,
	})
}
