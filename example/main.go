package main

import (
	"log"
	"net/http"

	gosimplessg "github.com/K-Sato1995/go-simple-ssg"

	"github.com/K-Sato1995/go-simple-ssg/config"
)

func main() {
	baseConfig := config.NewConfig(config.Config{
		SiteInfo: config.SiteInfo{
			Title:       "My custom Blog",
			Description: "This is my custom blog",
		},
	})
	engine := gosimplessg.New(baseConfig)
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
