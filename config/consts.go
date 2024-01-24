package config

const GENERATED_HTML_DIR = "generated"
const TEMPLATE_DIR = "templates"

type Config struct {
	TemplatePath        string // dir that contains base html/css
	GeneratedPath       string // dir that contains generated html/css
	HotReloadServerPort int
}

func NewConfig(custom Config) Config {
	config := Config{
		TemplatePath:        TEMPLATE_DIR,
		GeneratedPath:       GENERATED_HTML_DIR,
		HotReloadServerPort: 8080,
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
