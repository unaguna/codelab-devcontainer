package main

import (
	"html/template"
	"os"
)

func main() {
	type Codelab struct {
		Title string
		Id    string
	}
	type Model struct {
		Codelabs []Codelab
	}

	codelabs := []Codelab{
		{Title: "title1", Id: "id1"},
		{Title: "title2", Id: "id2"},
	}
	model := Model{Codelabs: codelabs}

	tmpl, err := template.ParseFiles("./src/index.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, model)
	if err != nil {
		panic(err)
	}
}
