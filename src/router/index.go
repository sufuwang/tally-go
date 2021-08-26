package router

import (
	"net/http"
	"tally-go/src/middleware"
	"tally-go/src/router/user"
)

// http.Handle("/register", Middleware.CORSMiddleWare(http.HandlerFunc(handleRegister)))

func RegisterRouter() {
	ds := user.RegisterUserRouter()
	for k, v := range ds {
		http.Handle(k, middleware.CORSMiddleWare(http.HandlerFunc(v)))
	}
}
