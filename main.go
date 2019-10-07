package main

import (
	"fmt"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

var (
	_url = "https://transit.yahoo.co.jp/traininfo/area/4/"
)

func main() {
	doc, err := goquery.NewDocument(_url)
	if err != nil {
		panic(err)
	}
	u := url.URL{}
	u.Scheme = doc.Url.Scheme
	u.Host = doc.Url.Host

	fmt.Println(u.Scheme)
	fmt.Println(u.Host)

	title := doc.Find("title").Text()
	fmt.Println(title)
}
