package engine

import (
	"go-simple-ssg/builder"
	"go-simple-ssg/config"
	"log"
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
