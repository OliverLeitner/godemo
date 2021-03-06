// +build ignore

package main

/**
* from:
* https://blog.gopheracademy.com/advent-2017/using-go-templates/
* https://golang.org/doc/articles/wiki/
 */

import (
	"fmt"
	"godemo/src/libs/grabmyown"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

/*func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}
}*/

/*func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}*/

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func namesToParagraphs() string {
	var data = grabmyown.CustomersAsArray()
	var listBody string
	for i := 0; i < len(data); i++ {
		listBody += "<p>" + data[i].ContactFirstName + "</p>"
	}
	return listBody
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// title := r.URL.Path[len("/customers"):]
	// p, _ := loadPage(title)
	/*var data = grabmyown.CustomersAsArray()
	var listBody string
	for i := 0; i < len(data); i++ {
		// fmt.Fprintf(w, "<h1>%s</h1><p>%s</p>", p.Title, data[i].ContactFirstName)
		listBody += "<p>" + data[i].ContactFirstName + "</p>"
	}*/
	// var bodyData = namesToParagraphs()
	t, _ := template.ParseFiles("./templates/main.html")
	t.Execute(w, grabmyown.CustomersAsArray())
	// fmt.Fprintf(w, "<div>%s</div>", namesToParagraphs())
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/customers/", viewHandler)
	log.Fatal(http.ListenAndServe("192.168.2.185:8080", nil))
}
