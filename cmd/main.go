package main

import (
	"github.com/dawsonalex/site-builder/tpl"
	"html/template"
	"log"
	"net/http"
)

func main() {

	//tpl, err := template.New("master").ParseFiles("/home/ad/GolandProjects/site-builder/templates/index.html", "/home/ad/GolandProjects/site-builder/templates/header.html")

	// TODO: blocks for components
	// TODO: build page from components
	// TODO: write components to XML
	// TODO: decode components from XML
	// TODO: How are we going to handle scope of components

	//dynamicTemplateFunc := func(name string, data any) error {
	//	return tpl.ExecuteTemplate(os.Stdout, name, data)
	//}
	//tpl.Funcs(map[string]interface{}{"dynamicTemplate": dynamicTemplateFunc})
	//tpl, err := template.New("master").ParseFiles("../templates/index.html")
	//handleErr(err)
	//err = tpl.ExecuteTemplate(os.Stdout, "hero", nil)
	//handleErr(err)
	////err = tpl.Execute(os.Stdout, map[string]string{"templateName": "header"})
	//handleErr(err)

	http.HandleFunc("/preview", func(writer http.ResponseWriter, request *http.Request) {
		t, err := template.New("master").ParseGlob("/home/ad/GolandProjects/site-builder/templates/*.html")
		handleErr(err)

		component := tpl.ComponentImpl[tpl.HeroConfig]{
			Conf: tpl.HeroConfig{HeroText: ""},
		}
		err = t.ExecuteTemplate(writer, component.Id, component.Config())

	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
