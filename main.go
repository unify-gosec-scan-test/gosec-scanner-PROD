package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	LoadEnv()

	http.HandleFunc("/api/login", LoginHandler)
	http.HandleFunc("/api/register", RegisterHandler)
	http.HandleFunc("/api/profile", AuthMiddleware(ProfileHandler))
	http.HandleFunc("/files/", FileServerHandler)

	log.Println("🚨 Starting vulnerable API on :8081")
	err := http.ListenAndServe(":8081", nil) // Insecure HTTP
	if err != nil {
		log.Fatal("Failed to start:", err)
	}
}
