package main

import ("fmt"; "net/http"; "io/ioutil"; "encoding/xml")

// SitemapIndex is a struct to store sitemap data
type SitemapIndex struct {
  // remember to let go of any whitespase in xml slice
  Locations []Location `xml:"sitemap"`
}

//Location is a slice inside sitemap, stores string between <loc> and </loc>
type Location struct{
  Loc string `xml:"loc"`
}
func (l Location) String() string {
  return fmt.Sprintf(l.Loc)
}

func main(){
  resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
  // resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/world.xml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  // str := string(bytes)
  // fmt.Println(str)
  resp.Body.Close()

  var s SitemapIndex
  xml.Unmarshal(bytes, &s)

  // fmt.Println(s)
  for _, Location := range s.Locations {
    fmt.Printf("\n%s", Location)
  }
}
