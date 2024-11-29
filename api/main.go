package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arinji2/garconia/sqlite"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})
	fmt.Println("Server Started")
	db, err := sqlite.NewConnection()
	fmt.Println("Connection Established")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("DB Found And Ready")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
