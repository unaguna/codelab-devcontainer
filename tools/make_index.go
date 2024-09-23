package main

import (
	"encoding/json"
	"html/template"
	"io/fs"
	"os"
	"path"
)

type Codelab struct {
	Title string
	Id    string
}

type Model struct {
	Codelabs []Codelab
}

func main() {

	distDir := "./dist"
	codelabFiles, err := fs.Glob(os.DirFS(distDir), "**/codelab.json")
	if err != nil {
		panic(err)
	}

	codelabs, err := parseCodelabJsons(codelabFiles, distDir)
	if err != nil {
		panic(err)
	}

	model := Model{Codelabs: codelabs}
	if err := outputIndexHtml(model, distDir); err != nil {
		panic(err)
	}
}

// 引数に指定した codelab.json ファイルをパースする。
func parseCodelabJsons(codelabJsonFiles []string, distDir string) ([]Codelab, error) {
	codelabs := make([]Codelab, 0, 10)

	for _, filepath := range codelabJsonFiles {
		file, err := os.Open(path.Join(distDir, filepath))
		if err != nil {
			return nil, err
		}

		var c Codelab
		if err := json.NewDecoder(file).Decode(&c); err != nil {
			return nil, err
		}

		codelabs = append(codelabs, c)
	}

	return codelabs, nil
}

// index.html を出力
func outputIndexHtml(model Model, distDir string) error {
	tmpl, err := template.ParseFiles("./src/index.html")
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path.Join(distDir, "index.html"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, model)
	if err != nil {
		return err
	}

	return nil
}
