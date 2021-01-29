package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
		appName := getEnv("APP_NAME", "")
		if appName != "" {
			fmt.Fprintf(w, "\nApp name: %s", appName)
		}
	})
	http.ListenAndServe(":"+getEnv("APP_PORT", "8080"), nil)
}

func getEnv(key, defaultVal string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultVal
	}
	return value
}
