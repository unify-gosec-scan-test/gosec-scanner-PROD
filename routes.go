package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = make(map[string]string) // username:password (plaintext)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var u User
	_ = json.NewDecoder(r.Body).Decode(&u)

	// No validation or password hashing
	users[u.Username] = u.Password

	log.Printf("📥 Registered user: %s with password: %s", u.Username, u.Password)
	fmt.Fprintf(w, "User %s registered", u.Username)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var u User
	_ = json.NewDecoder(r.Body).Decode(&u)

	// Log sensitive info
	log.Printf("🔐 Login attempt: %s/%s", u.Username, u.Password)

	if pass, ok := users[u.Username]; ok && pass == u.Password {
		token := GenerateToken(u.Username)
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	} else {
		http.Error(w, "Invalid credentials", 401)
	}
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("X-User")
	fmt.Fprintf(w, "Welcome to your profile, %s!", user)
}

func FileServerHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/files/"):]
	data, err := io.ReadFile("./static/" + path) // 🧨 vulnerable to ../ traversal
	if err != nil {
		http.Error(w, "File not found: "+err.Error(), 404)
		return
	}
	w.Write(data)
}
