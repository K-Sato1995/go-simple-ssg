package builder

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/K-Sato1995/go-simple-ssg/config"
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

func GenerateListPage(articles []ArticleInfo, templatePath, generatedPath string, siteInfo config.SiteInfo) error {
	listTmpl, err := template.ParseFiles(filepath.Join(templatePath, "list.html"))
	if err != nil {
		return err
	}
	if _, err := os.Stat(generatedPath); os.IsNotExist(err) {
		os.Mkdir(generatedPath, 0755)
	}

	data := ListPageData{
		HTMLTitle:       siteInfo.Title,
		MetaDescription: siteInfo.Description,
		PageTitle:       siteInfo.Title,
		Articles:        articles,
	}

	var renderedContent bytes.Buffer
	err = listTmpl.Execute(&renderedContent, data)
	if err != nil {
		log.Fatal("Error executing list template:", err)
		return err
	}
	ioutil.WriteFile(filepath.Join(generatedPath, "index.html"), renderedContent.Bytes(), 0644)
	return nil
}
