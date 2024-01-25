package config

const GENERATED_HTML_DIR = "generated"
const TEMPLATE_DIR = "templates"

type SiteInfo struct {
	Title       string
	Description string
}
type Config struct {
	TemplatePath        string // dir that contains base html/css
	GeneratedPath       string // dir that contains generated html/css
	SiteInfo            SiteInfo
	HotReloadServerPort int
}

func NewConfig(custom Config) Config {
	config := Config{
		TemplatePath:  TEMPLATE_DIR,
		GeneratedPath: GENERATED_HTML_DIR,
		SiteInfo: SiteInfo{
			Title:       "My Blog",
			Description: "This is my blog",
		},
		HotReloadServerPort: 8080,
	}
	if custom.SiteInfo.Title != "" {
		config.SiteInfo.Title = custom.SiteInfo.Title
	}
	if custom.SiteInfo.Description != "" {
		config.SiteInfo.Description = custom.SiteInfo.Description
	}
	if custom.TemplatePath != "" {
		config.TemplatePath = custom.TemplatePath
	}
	if custom.GeneratedPath != "" {
		config.GeneratedPath = custom.GeneratedPath
	}
	if custom.HotReloadServerPort != 0 {
		config.HotReloadServerPort = custom.HotReloadServerPort
	}
	return config
}
