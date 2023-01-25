package web

import (
	"log"
	"text/template"
)

type Web struct {
	Template *template.Template
}

func NewWeb(files ...string) *Web {
	files = append(files, "web/pages/header.gohtml", "web/pages/menu.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
	}
	return &Web{
		Template: t,
	}
}
