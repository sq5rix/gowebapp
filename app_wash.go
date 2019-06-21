package main

import ("fmt"; "net/http"; "io/ioutil"; "encoding/xml"; "log"; "strings")

// SitemapIndex is a struct to store sitemap string
type SitemapIndex struct {
  // remember to let go of any whitespase in xml slice
  Locations []string `xml:"sitemap>loc"`
}

// News - store location from sitemap
type News struct {
  Articles []string `xml:"url>loc"`
}

func main(){
  var s SitemapIndex
  var n News
  // nmap := make(map[string] NewsMap)

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

  for _, loc := range s.Locations {
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
    for _, art := range n.Articles {
      fmt.Println("\n\n")
      fmt.Println(art)
    }
  }

}
