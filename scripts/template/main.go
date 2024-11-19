package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

//go:embed tmpls/*
var fs embed.FS

func main() {
	today := time.Now()
	day := flag.Int("day", today.Day(), "day number to fetch, 0-25")
	flag.Parse()
	Run(*day)
}

// Run makes a template main.go and main_test.go file for the given day
func Run(day int) {
	if day > 25 || day < 0 {
		log.Fatalf("invalid -day value, must be 0 through 25, got %v", day)
	}

	ts, errTemplate := template.ParseFS(fs, "tmpls/*")
	if errTemplate != nil {
		log.Fatalf("parsing template directory: %s", errTemplate)
	}

	mainFilename := fmt.Sprintf("day-%d/main.go", day)
	testFilename := fmt.Sprintf("day-%d/main_test.go", day)
	inputFilename := fmt.Sprintf("day-%d/input.txt", day)
	readmeFilename := fmt.Sprintf("day-%d/README.md", day)

	mkDirErr := os.MkdirAll(filepath.Dir(mainFilename), os.ModePerm)
	if mkDirErr != nil {
		log.Fatalf("creating base dir: %s", mkDirErr)
	}

	ensureNotOverwriting(mainFilename)
	ensureNotOverwriting(testFilename)
	ensureNotOverwriting(inputFilename)
	ensureNotOverwriting(readmeFilename)

	mainFile, errMain := os.Create(mainFilename)
	if errMain != nil {
		log.Fatalf("creating main.go file: %v", errMain)
	}

	testFile, errTest := os.Create(testFilename)
	if errTest != nil {
		log.Fatalf("creating main_test.go file: %v", errTest)
	}

	_, errInput := os.Create(inputFilename)
	if errInput != nil {
		log.Fatalf("creating input.txt file: %v", errInput)
	}

	readmeFile, errReadme := os.Create(readmeFilename)
	if errReadme != nil {
		log.Fatalf("creating Readme.md file: %v", errReadme)
	}

	templateData := struct{ Day int }{day}

	mainTemplateErr := ts.ExecuteTemplate(mainFile, "main.go.tmpl", templateData)
	if mainTemplateErr != nil {
		log.Fatalf("main template: %s", mainTemplateErr)
	}
	mainTestTemplateErr := ts.ExecuteTemplate(testFile, "main_test.go.tmpl", templateData)
	if mainTestTemplateErr != nil {
		log.Fatalf("main test template: %s", mainTestTemplateErr)
	}
	readmeTemplateErr := ts.ExecuteTemplate(readmeFile, "README.md.tmpl", templateData)
	if readmeTemplateErr != nil {
		log.Fatalf("main test template: %s", readmeTemplateErr)
	}
	fmt.Printf("templates made for day-%d\n", day)
}

func ensureNotOverwriting(filename string) {
	_, err := os.Stat(filename)
	if err == nil {
		log.Fatalf("File already exists: %s", filename)
	}
}
