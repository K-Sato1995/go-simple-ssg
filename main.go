package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

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

func mdToHTML(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	return markdown.Render(doc, renderer)
}

func main() {
	// content, err := ioutil.ReadFile("./contents/testMd.md")
	// if err != nil {
	// 	log.Fatalf("Error reading Markdown file: %v", err)
	// }
	// md := []byte(content)
	// html := mdToHTML(md)
	// fmt.Println(string(html))
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
			data := Template{
				HTMLTitle:       "Example Title",
				MetaDescription: "Example Description",
				PageTitle:       "My Page Title",
				Content:         template.HTML(content),
			}

			var renderedContent bytes.Buffer
			err = htmlTmpl.Execute(&renderedContent, data)
			if err != nil {
				log.Println("Error executing template:", err)
				return nil
			}
			fmt.Println(renderedContent.String())
		}
		return nil
	})
}

func fileWalk() {

}
