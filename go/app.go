package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"text/template"
	"time"
)

type Params struct {
	Name string
	Time string
}

func handler(w http.ResponseWriter, r *http.Request) {
	const layout = "3:04pm"

	re := regexp.MustCompile("\\w+")
	name := re.FindString(r.URL.RequestURI())

	html, _ := ioutil.ReadFile("./index.html")

	now := time.Now()
	params := Params{name, now.Format(layout)}

	t, _ := template.New("index").Parse(string(html))
	t.Execute(w, params)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
