package main

import (
	"log"
	"net/http"
	"text/template"
	"path/filepath"
	"sync"
)

//temp1 represents a single template
type templateHandler struct {
	once		sync.Once
	filename	string
	temp1		*template.Template
}

//ServeHTTP handles the HTTP Request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	t.once.Do(func() {
		t.temp1 = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
		})
	t.temp1.Execute(w, nil)
}


func main() {
	r := newRoom()

	//root
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	//get the room going
	go r.run()
	//start the web server 
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}