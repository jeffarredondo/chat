package main

import (
	"log"
	"net/http"
	"text/template"
	"path/filepath"
	"sync"
	"flag"
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
	t.temp1.Execute(w, r)
}


func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	r := newRoom()
	flag.Parse() //parse the flags
	//root
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	//get the room going
	go r.run()
	//start the web server 
	log.Println("Starting the web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}