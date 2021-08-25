package tool

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

// test=Test_Cookie;SameSite=None;Secure;Max-Age=10;
type TypeCookieOption struct {
	Key      string
	Value    string
	SameSite string
	Secure   string
	Expires  string
	MaxAge   string
}

func defaultOption(key string, value string) TypeCookieOption {
	option := TypeCookieOption{
		Key:      key,
		Value:    value,
		SameSite: "None",                   // 支持跨域
		Secure:   "false",                  // 非 HTTPS
		MaxAge:   fmt.Sprint(24 * 60 * 60), // 有效时间 1 Day
	}
	a := reflect.TypeOf('1').String()
	fmt.Println(a)
	return option
}

func cookieToString(cookie TypeCookieOption) string {
	ds := reflect.ValueOf(cookie)
	var (
		cookieStr       string
		cookieOptionStr string
		key             string
		value           string
	)
	for i := 0; i < ds.NumField(); i++ {
		key = ds.Type().Field(i).Name
		value = ds.Field(i).String()
		fmt.Printf("Key: %s, %s\n", key, value)
		if len(value) > 0 {
			if key == "Key" {
				cookieStr = value + cookieStr
			} else if key == "Value" {
				cookieStr += "=" + value + ";"
			} else if key == "Secure" && value == "true" {
				cookieOptionStr += "Secure;"
			} else if key == "MaxAge" {
				cookieOptionStr += fmt.Sprintf("%s=%s;", "Max-Age", value)
			} else {
				cookieOptionStr += fmt.Sprintf("%s=%s;", key, value)
			}
		}
	}
	return cookieStr + cookieOptionStr
}

func SetCookieToHeader(w http.ResponseWriter, key string, value string) {
	cookie := defaultOption(key, value)
	cookieStr := cookieToString(cookie)
	w.Header().Set("Set-Cookie", cookieStr)
}

func GetSingleCookieFromCookies(cookies string, key string) string {
	if len(cookies) == 0 || !strings.Contains(cookies, key) {
		return ""
	}
	slice := strings.Split(cookies, "; ")
	var cookie string
	var str string
	for index := range slice {
		str = slice[index]
		if strings.Contains(str, key+"=") {
			cookie = str
		}
	}
	return strings.Split(cookie, "=")[1]
}
