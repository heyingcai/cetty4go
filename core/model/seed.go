package model

import "net/http"

type Seed struct {
	Url      string
	Method   string
	Postbody string
	Cookies  []*http.Cookie
	Headers  []*http.Header
	Attach   map[string]string
}

func NewGetSeed(url string) {

}

func NewSeed(url string, method string, cookies []*http.Cookie, headers []*http.Header) *Seed {
	return &Seed{url, method, "", cookies, headers, make(map[string]string, 10)}
}

func (seed *Seed) GetUrl() string {
	return seed.Url
}

func (seed *Seed) GetMethod() string {
	return seed.Method
}

func (seed *Seed) GetPostbody() string {
	return seed.Postbody
}

func (seed *Seed) GetCookies() []*http.Cookie {
	return seed.Cookies
}

func (seed *Seed) GetHeader() []*http.Header {
	return seed.Headers
}

func (seed *Seed) GetAttach() map[string]string {
	return seed.Attach
}
