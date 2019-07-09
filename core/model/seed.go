package model

import "net/http"

type Seed struct {
	url    string
	method string
	cookies []*http.Cookie
	headers []*http.Header
	attach map[string]string


}
