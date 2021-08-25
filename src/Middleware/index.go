package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

var FrontEndUrl = "http://127.0.0.1:8080"

// 跨域校验: 为符合跨域规则的请求响应加上 Access-Control-Allow-Origin
func CORSMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host := r.Host
		isPassCORS := strings.Contains(host, "localhost") || strings.Contains(host, "127.0.0.1")
		fmt.Println("isPassCORS: ", isPassCORS)
		if isPassCORS {
			w.Header().Set("Access-Control-Allow-Origin", FrontEndUrl)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			next.ServeHTTP(w, r)
		}
	})
}
