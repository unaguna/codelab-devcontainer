//usr/local/go/bin/go run $0 $@ ; exit

package main

import (
	"cmp"
	"encoding/json"
	"flag"
	"html/template"
	"io/fs"
	"os"
	"path"
	"slices"
)

type Codelab struct {
	Title string
	Id    string
}

type Model struct {
	Codelabs []Codelab
}

func main() {
	flag.Parse()
	var codelabDirs = flag.Args()

	indexSrcPath, ok := os.LookupEnv("INDEX_SRC_PATH")
	if !ok {
		indexSrcPath = "./src/index.html"
	}

	distDir, ok := os.LookupEnv("DIST_DIR")
	if !ok {
		distDir = "/workspace_local/dist"
	}

	if len(codelabDirs) <= 0 {
		codelabDirs = append(codelabDirs, distDir)
	}

	var codelabFiles = make([]string, 0, 10)
	for _, codelabDir := range codelabDirs {
		files, err := fs.Glob(os.DirFS(codelabDir), "**/codelab.json")
		if err != nil {
			panic(err)
		}
		for i, file := range files {
			files[i] = path.Join(codelabDir, file)
		}
		codelabFiles = append(codelabFiles, files...)
	}

	codelabs, err := parseCodelabJsons(codelabFiles)
	if err != nil {
		panic(err)
	}

	model := Model{Codelabs: codelabs}
	if err := outputIndexHtml(model, distDir, indexSrcPath); err != nil {
		panic(err)
	}
}

// 引数に指定した codelab.json ファイルをパースする。
func parseCodelabJsons(codelabJsonFiles []string) ([]Codelab, error) {
	codelabs := make([]Codelab, 0, 10)

	for _, filepath := range codelabJsonFiles {
		file, err := os.Open(filepath)
		if err != nil {
			return nil, err
		}

		var c Codelab
		if err := json.NewDecoder(file).Decode(&c); err != nil {
			return nil, err
		}

		codelabs = append(codelabs, c)
	}

	// Distinct by Id
	slices.SortFunc(codelabs, func(a, b Codelab) int { return cmp.Compare(a.Id, b.Id) })
	codelabs = slices.CompactFunc(codelabs, func(a, b Codelab) bool { return a.Id == b.Id })

	return codelabs, nil
}

// index.html を出力
func outputIndexHtml(model Model, distDir string, srcPath string) error {
	tmpl, err := template.ParseFiles(srcPath)
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
