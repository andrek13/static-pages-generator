package main

import (
	"flag"
	"fmt"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	buildCmd := flag.NewFlagSet("build", flag.ExitOnError)
	serveCmd := flag.NewFlagSet("serve", flag.ExitOnError)

	buildInputDir := buildCmd.String("input", "content", "Directory containing Markdown files")
	buildOutputDir := buildCmd.String("output", "public", "Directory to save HTML files")
	servePort := serveCmd.String("port", "8080", "Port for the HTTP server")

	if len(os.Args) < 2 {
		fmt.Println("Usage: static-pages-generator [build|serve] [options]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "build":
		buildCmd.Parse(os.Args[2:])
		build(*buildInputDir, *buildOutputDir)
	case "serve":
		serveCmd.Parse(os.Args[2:])
		serve(*servePort)
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		fmt.Println("Usage: static-pages-generator [build|serve] [options]")
		os.Exit(1)
	}
}

func build(inputDir, outputDir string) {
	fmt.Printf("Building from %s to %s\n", inputDir, outputDir)
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

		html := blackfriday.Run(data)
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
		fmt.Printf("Error during build: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Build complete.")
}

func serve(port string) {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	fmt.Printf("Serving on :%s\n", port)

	server := &http.Server{
		Addr: ":" + port,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error serving: %v\n", err)
		os.Exit(1)
	}
}
