package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
		// second value is data
		t.templ.Execute(w, r)
	})
}

func main() {
	var p = flag.String("port", ":8080", "address of app")
	flag.Parse()

	http.Handle("/", &templateHandler{filename: "chat.html"})

	r := newRoom()
	//r.tracer = trace.New(os.Stdout)
	http.Handle("/room", r)

	go r.run()

	log.Println("launch server on port:", *p)

	if err := http.ListenAndServe(*p, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
