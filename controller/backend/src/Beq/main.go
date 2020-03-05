package main

import (
	"log"
	"net/http"
	"os"

	routes "Beq/routes"

	"github.com/joho/godotenv"
)

func main() {

	e := godotenv.Load()

	if e != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	// // Handle routes
	http.Handle("/", routes.Handlers())

	// // serve
	log.Println("INFO: Server Online @Port:" + port)
	log.Println("INFO: URL :" + "http://localhost:" + port + "/")

	// log.Fatal(http.ListenAndServeTLS(":"+port, "server.crt", "server.key", nil))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
