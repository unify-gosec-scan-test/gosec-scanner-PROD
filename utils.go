package main

import (
	"log"
	"os"
)

func LoadEnv() {
	file, err := os.Open(".env")
	if err != nil {
		log.Println("⚠️ No .env found, continuing anyway")
		return
	}
	defer file.Close()

	var key, val string
	for {
		_, err := fmt.Fscanf(file, "%s=%s\n", &key, &val)
		if err != nil {
			break
		}
		os.Setenv(key, val)
	}
}
