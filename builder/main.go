package builder

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"site-generator/config"
	"site-generator/parser"
	"strings"
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

func GenerateListPage(articles []ArticleInfo) error {
	listTmpl, err := template.ParseFiles(filepath.Join(config.ASSETS_DIR, "list.html"))
	if err != nil {
		return err
	}
	if _, err := os.Stat(config.GENERATED_HTML_DIR); os.IsNotExist(err) {
		os.Mkdir(config.GENERATED_HTML_DIR, 0755)
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
		return err
	}
	ioutil.WriteFile(filepath.Join(config.GENERATED_HTML_DIR, "list.html"), renderedContent.Bytes(), 0644)
	return nil
}

func GenerateDetailPages() ([]ArticleInfo, error) {
	var articles []ArticleInfo
	htmlTmpl, err := template.ParseGlob(filepath.Join(config.ASSETS_DIR, "detail.html"))
	if err != nil {
		return nil, err
	}
	filepath.Walk("contents", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println("error reading file", err)
				return err
			}
			metadata, mdContent, err := parser.ParseMetadata(content)
			if err != nil {
				fmt.Println("error parsing metadata", err)
				return err
			}
			fmt.Println("meta", metadata)
			// List page==begin
			baseName := strings.TrimSuffix(filepath.Base(path), ".md")
			outputPath := filepath.Join(config.GENERATED_HTML_DIR, baseName+".html")
			articles = append(articles, ArticleInfo{
				Title: metadata.Title,
				Path:  outputPath,
			})
			// List page==end

			// Create detail page template
			parsedHtml := parser.MdToHTML(mdContent)
			data := Template{
				HTMLTitle:       "Example Title",
				MetaDescription: "Example Description",
				PageTitle:       "My Page Title",
				Content:         template.HTML(parsedHtml),
			}

			var renderedContent bytes.Buffer
			err = htmlTmpl.Execute(&renderedContent, data)
			if err != nil {
				fmt.Println("error executing template", err)
				return err
			}
			err = ioutil.WriteFile(outputPath, renderedContent.Bytes(), 0644)
			if err != nil {
				fmt.Println("error writing file", err)
				return err
			}
			fmt.Printf("Article written to %s\n", outputPath)
		}
		return nil
	})
	return articles, nil
}
