package main

import ("fmt"; "net/http")

func handler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, `<h1>hello go!!</h1>
                  <p>fun language</p> `)
}

func main(){
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8000", nil)
}
