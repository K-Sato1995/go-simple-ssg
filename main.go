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
			metadata, mdContent, err := parseMetadata(content)
			if err != nil {
				log.Println(err)
				return nil
			}
			fmt.Println("meta", metadata)
			data := Template{
				HTMLTitle:       "Example Title",
				MetaDescription: "Example Description",
				PageTitle:       "My Page Title",
				Content:         template.HTML(mdContent),
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
