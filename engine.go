package engine

import (
	"log"

	"github.com/K-Sato1995/go-simple-ssg/builder"
	"github.com/K-Sato1995/go-simple-ssg/config"
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
