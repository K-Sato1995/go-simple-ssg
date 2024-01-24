package main

import (
	"log"
	"net/http"
	"site-generator/builder"
	"site-generator/config"
)

func main() {
	builder.BundleAndMinifyCSS()
	// Create detail pages
	articles, err := builder.GenerateDetailPages()
	if err != nil {
		log.Fatal(`error occured while generateing detail pages`, err)
	}
	// Create list page
	builder.GenerateListPage(articles)

	// Serve files for development
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
