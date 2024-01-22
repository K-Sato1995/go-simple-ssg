package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Template struct {
	HTMLTitle       string
	MetaDescription string
	PageTitle       string
	Content         template.HTML
}

type ArticleInfo struct {
	Title string
	Path  string
}

type ListPageData struct {
	HTMLTitle       string
	MetaDescription string
	PageTitle       string
	Articles        []ArticleInfo
}

type MetaData struct {
	Title       string
	Description string
}

func mdToHTML(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	return markdown.Render(doc, renderer)
}

func parseMetadata(content []byte) (MetaData, []byte, error) {
	var meta MetaData
	sections := bytes.SplitN(content, []byte("----"), 3)
	if len(sections) < 3 {
		return meta, nil, errors.New("invalid format: metadata not found")
	}
	metadataContent := sections[1]
	mdContent := sections[2]
	lines := bytes.Split(metadataContent, []byte("\n"))
	for _, line := range lines {
		line = bytes.TrimLeft(line, "- ")
		keyValue := bytes.SplitN(line, []byte(":"), 2)
		if len(keyValue) != 2 {
			continue
		}
		key := string(bytes.TrimSpace(keyValue[0]))
		value := string(bytes.TrimSpace(keyValue[1]))
		value = strings.Trim(value, "\"")
		switch key {
		case "Title":
			meta.Title = value
		case "Description":
			meta.Description = value
		}
	}
	return meta, mdContent, nil
}

func generateListPage(articles []ArticleInfo) {
	listTmpl, err := template.ParseFiles("templates/list.html")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := os.Stat("public"); os.IsNotExist(err) {
		os.Mkdir("public", 0755)
	}

	data := ListPageData{
		HTMLTitle:       "Articles List",
		MetaDescription: "List of articles",
		PageTitle:       "Articles",
		Articles:        articles,
	}

	var renderedContent bytes.Buffer
	err = listTmpl.Execute(&renderedContent, data)
	if err != nil {
		log.Fatal("Error executing list template:", err)
	}
	ioutil.WriteFile("public/list.html", renderedContent.Bytes(), 0644)
}

func main() {
	var articles []ArticleInfo
	htmlTmpl, err := template.ParseGlob("templates/detail.html")

	if err != nil {
		log.Fatal(err)
	}
	filepath.Walk("contents", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == ".md" {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				log.Println(err)
				return nil
			}
			metadata, mdContent, err := parseMetadata(content)
			if err != nil {
				log.Println(err)
				return nil
			}
			fmt.Println("meta", metadata)
			// List page==begin
			baseName := strings.TrimSuffix(filepath.Base(path), ".md")
			outputPath := filepath.Join("public", baseName+".html")
			articles = append(articles, ArticleInfo{
				Title: metadata.Title,
				Path:  outputPath,
			})
			// List page==end

			// Create detail page template
			parsedHtml := mdToHTML(mdContent)
			data := Template{
				HTMLTitle:       "Example Title",
				MetaDescription: "Example Description",
				PageTitle:       "My Page Title",
				Content:         template.HTML(parsedHtml),
			}

			var renderedContent bytes.Buffer
			err = htmlTmpl.Execute(&renderedContent, data)
			if err != nil {
				log.Println("Error executing template:", err)
				return nil
			}
			err = ioutil.WriteFile(outputPath, renderedContent.Bytes(), 0644)
			if err != nil {
				log.Printf("Error writing file %s: %v", outputPath, err)
				return nil
			}
			fmt.Printf("Article written to %s\n", outputPath)
		}
		// Create list page
		generateListPage(articles)
		return nil
	})
}

func fileWalk() {

}
