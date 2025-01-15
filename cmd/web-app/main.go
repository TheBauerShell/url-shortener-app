package main

import (
	"database/sql"
	"github.com/TheBauerShell/url-shortener-app/controllers"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	slite, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatal(err)

	}
	defer slite.Close()

	http.HandleFunc("/", controllers.ShowIndex)
	http.HandleFunc("/shorten", controllers.Shorten(slite))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
