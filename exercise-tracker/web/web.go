package web

import (
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}
