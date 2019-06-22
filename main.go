package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// SitemapIndex main sitemap
type SitemapIndex struct {
	// remember to let go of any whitespase in xml slice
	Locations []string `xml:"sitemap>loc"`
}

// News - store location from sitemap
type News struct {
	Articles []string `xml:"url>loc"`
}

// news will go here
var n News

// Handler shows the Home Page
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>hello go!!</h1>
                  <p>fun language</p> `)
}

// NewsAggPage for template
type NewsAggPage struct {
	Title string
	News  News
}

func aggHandler(w http.ResponseWriter, r *http.Request) {

	p := NewsAggPage{Title: "News Handler"}
	p.News = n
	fmt.Println(p.News)
	t, _ := template.ParseFiles("basic.html")
	t.Execute(w, p)
}

func main() {

	var s SitemapIndex

	resp, err := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
	if err != nil {
		log.Fatalln(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	xml.Unmarshal(bytes, &s)

	defer resp.Body.Close()

	for idx, loc := range s.Locations {
		// added strings.TrimSpace after error
		//: net/url: invalid control character in URL
		// it was \r
		resp, err := http.Get(strings.TrimSpace(loc))
		if err != nil {
			log.Fatalln(err)
		}
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		xml.Unmarshal(bytes, &n)
		if idx >= 5 {
			break
		}
	}
	http.HandleFunc("/", Handler)
	http.HandleFunc("/agg/", aggHandler)
	http.ListenAndServe(":8888", nil)
}
