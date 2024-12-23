package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: static-pages-generator dashboard-template.md")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "build":
		build()
	case "serve":
		serve()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func build() {
	inputDir := "content"
	outputDir := "public"

	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		html := blackfriday.MarkdownCommon(data)
		outputPath := filepath.Join(outputDir, strings.Replace(info.Name(), ".md", ".html", 1))
		err = os.MkdirAll(filepath.Dir(outputPath), os.ModePerm)
		if err != nil {
			return err
		}

		tmpl, err := template.New("page").Parse(`<html><body>{{.Content}}</body></html>`)
		if err != nil {
			return err
		}

		file, err := os.Create(outputPath)
		if err != nil {
			return err
		}
		defer file.Close()

		err = tmpl.Execute(file, map[string]interface{}{
			"Content": template.HTML(html),
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func serve() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	fmt.Println("Serving on :8080")

	server := &http.Server{
		Addr: ":8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error serving: %v\n", err)
		os.Exit(1)
	}
}
