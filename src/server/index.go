package server

import (
	"fmt"
	"net/http"
)

type StartParameter struct {
	Domain    string
	HttpPort  int
	HttpsPort int
}

func Start(P StartParameter) {
	domain := P.Domain
	httpPort := P.HttpPort
	httpsPort := P.HttpsPort

	str := "HTTPS: 当前模式:"
	isDevelopment := domain != "sufuwang.top"
	if isDevelopment {
		str += "开发"
	} else {
		str += "生产"
	}
	fmt.Println(str)

	if httpPort != -1 && httpsPort == -1 {
		fmt.Println("命中开启方式 1")
		startHTTP(domain, httpPort)
	} else if httpPort == -1 && httpsPort != -1 {
		fmt.Println("命中开启方式 2")
		startHTTPS(domain, httpsPort, isDevelopment)
	} else if httpPort != -1 && httpsPort != -1 {
		fmt.Println("命中开启方式 3")
		go startHTTPS(domain, httpsPort, isDevelopment)
		startHTTP(domain, httpPort)
	}
}

func startHTTPS(domain string, port int, isDevelopment bool) {
	url := fmt.Sprintf("%s:%d", domain, port)
	var err error
	// "SSL证书公匙文件名"，"SSL证书密匙文件名"
	if isDevelopment {
		err = http.ListenAndServeTLS(
			url, "src/public/credential/cert.pem", "src/public/credential/key.pem", nil)
		// err = http.ListenAndServeTLS(
		// 	url, "src/public/credential/ca_certificate.pem", "src/public/credential/ca_key.pem", nil)
	} else {
		err = http.ListenAndServeTLS(
			url, "src/public/credential/1_sufuwang.top_bundle.crt", "src/public/credential/2_sufuwang.top.key", nil)
	}
	if err != nil {
		fmt.Println("HTTPS: ", err)
	}
}

func startHTTP(domain string, port int) {
	url := fmt.Sprintf("%s:%d", domain, port)
	err := http.ListenAndServe(url, nil)
	if err != nil {
		fmt.Println("HTTP: ", err)
	}
}
