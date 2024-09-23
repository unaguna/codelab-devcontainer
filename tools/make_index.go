package main

import (
	"encoding/json"
	"html/template"
	"io/fs"
	"os"
	"path"
)

func main() {
	type Codelab struct {
		Title string
		Id    string
	}
	type Model struct {
		Codelabs []Codelab
	}

	codelabs := make([]Codelab, 0, 10)

	distDir := "./dist"
	codelabFiles, err := fs.Glob(os.DirFS(distDir), "**/codelab.json")
	if err != nil {
		panic(err)
	}

	for _, filepath := range codelabFiles {
		file, err := os.Open(path.Join(distDir, filepath))
		if err != nil {
			panic(err)
		}

		var c Codelab
		if err := json.NewDecoder(file).Decode(&c); err != nil {
			panic(err)
		}

		codelabs = append(codelabs, c)
	}

	model := Model{Codelabs: codelabs}

	tmpl, err := template.ParseFiles("./src/index.html")
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(path.Join(distDir, "index.html"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(file, model)
	if err != nil {
		panic(err)
	}
}
