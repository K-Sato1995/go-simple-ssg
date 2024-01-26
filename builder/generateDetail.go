package builder

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/K-Sato1995/go-simple-ssg/config"
	"github.com/K-Sato1995/go-simple-ssg/parser"
)

func GenerateDetailPages(templatePath string, generatedPath string, siteInfo config.SiteInfo) ([]ArticleInfo, error) {
	var articles []ArticleInfo
	htmlTmpl, err := template.ParseGlob(filepath.Join(templatePath, "detail.html"))
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
			baseName := strings.TrimSuffix(filepath.Base(path), ".md")
			outputPath := filepath.Join(generatedPath, baseName+".html")
			articles = append(articles, ArticleInfo{
				Title:         metadata.Title,
				Description:   metadata.Description,
				PublishedDate: metadata.PublishedDate,
				Path:          baseName + ".html",
			})

			// Create detail page template
			parsedHtml := parser.MdToHTML(mdContent)
			data := Template{
				HTMLTitle:       siteInfo.Title,
				MetaDescription: metadata.Description,
				PageTitle:       metadata.Title,
				PublishedDate:   metadata.PublishedDate,
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
			// fmt.Printf("Article written to %s\n", outputPath)
		}
		return nil
	})
	return articles, nil
}
