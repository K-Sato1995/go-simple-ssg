package main

import (
	gosimplessg "go-simple-ssg"
	"go-simple-ssg/config"
	"log"
	"net/http"
)

func main() {
	baseConfig := config.NewConfig(config.Config{})
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
