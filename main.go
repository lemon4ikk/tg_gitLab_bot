package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	//tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	values, err := url.ParseQuery(string(body))
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(values)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/gitLabBot", handler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("error listening to the port: %v", err)
	}
}
