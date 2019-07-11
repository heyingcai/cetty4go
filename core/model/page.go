package model

import "github.com/PuerkitoBio/goquery"

type Page struct {
	Url       string
	Seed      *Seed
	Result    *Result
	Body      string
	docParser *goquery.Document
	NextSeeds []*Seed
}
