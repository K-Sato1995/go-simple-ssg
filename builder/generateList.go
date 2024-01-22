package builder

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"site-generator/config"
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
