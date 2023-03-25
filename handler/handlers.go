package handler

import (
	"fmt"
	"mercadofoto/app"
	"net/http"
	"text/template"
)

var NAME string = "MercadoFoto"
var URL string = "http://localhost:3000/"
var THEME_COLOR = "#598cff"

type Data struct {
	Name        string
	Url         string
	ThemeColor  string
	Title       string
	Description string
	Thumbnail   string
}

// DEFAULT SETTING FOR HTML PAGE
var INIT_SETTING Data = Data{
	Name:       NAME,
	Url:        URL,
	ThemeColor: THEME_COLOR,
}

func HomePage(rw http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(app.RenderLayout("home")...)

	if err != nil {
		panic(err)
	}

	INIT_SETTING := Data{
		Title:       "este es el titulo",
		Description: "Description of page",
	}

	t.Execute(rw, INIT_SETTING)
	// fmt.Fprintln(rw, "este es el home page")
}

func Landing(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "landing...")
}

func GetTest(rw http.ResponseWriter, r *http.Request) {
	Run(rw, r)
}
