package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
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

// NewsMap index it=s title, body is article
// type NewsMap map[string]string

// NewsStruct to pass to template
type NewsStruct struct {
	Title string
	News  map[string]string
}

// Data pass to template
var Data NewsStruct

// Handler shows the Home Page
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>hello go!!</h1>
                  <p>fun language</p> `)
}

// FindTitle of article in the url
func FindTitle(s string) string {
	re := regexp.MustCompile(`([a-z0-9]+-)+[a-z0-9]+`)
	s = re.FindString(s)
	s = strings.ReplaceAll(s, "-", " ")
	return strings.Title(s)
}

func aggHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(Data)
	t, _ := template.ParseFiles("basic.html")
	t.Execute(w, Data)
}

func main() {

	var s SitemapIndex
	var n News

	Data.Title = "Titles"
	Data.News = make(map[string]string)

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
		for _, v := range n.Articles {
			i := FindTitle(string(v))
			fmt.Println(i)
			Data.News[i] = v
		}

		if idx >= 1 {
			break
		}
	}
	http.HandleFunc("/", Handler)
	http.HandleFunc("/agg/", aggHandler)
	http.ListenAndServe(":8888", nil)
}
