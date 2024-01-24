package main

import (
	"log"
	"net/http"
	"site-generator/builder"
	"site-generator/config"
)

type Engine struct {
	Config *config.Config
	// Logger                  *Logger
	// HotReload               *HotReload
}

func New(config config.Config) *Engine {
	return &Engine{
		Config: &config,
	}
}

func (e *Engine) Build() {
	// Bundle css
	err := builder.BundleCSS(e.Config.TemplatePath, e.Config.GeneratedPath)
	if err != nil {
		log.Fatal(`error occured while bundling css`, err)
	}
	// Create detail pages
	articles, err := builder.GenerateDetailPages(e.Config.TemplatePath, e.Config.GeneratedPath)
	if err != nil {
		log.Fatal(`error occured while generateing detail pages`, err)
	}
	// Create list page
	builder.GenerateListPage(articles, e.Config.TemplatePath, e.Config.GeneratedPath)
}

func main() {
	baseConfig := config.NewConfig(config.Config{})
	engine := New(baseConfig)
	engine.Build()
	serveFiles()

}

func serveFiles() {
	fs := http.FileServer(http.Dir(config.GENERATED_HTML_DIR))
	http.Handle("/", fs)

	log.Println("Serving files on http://localhost:8080...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
