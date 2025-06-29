package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	//name := data["name"].(string)
	project := data["project"].(string)
	//event := data["event_type"].(string)

	bot, err := tgbotapi.NewBotAPI("7653357171:AAH0irLiLcA8TvZTXHNY3f_pCZMzVPMN0_E")
	chatID := int64(414747434) // ID бота
	//chatID := int64(-1002561300903) // ID тестового чата
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	//m := string(body)
	m := string(project)
	msg := tgbotapi.NewMessage(chatID, m)
	bot.Send(msg)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(body)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /gitLabBot", handler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("error listening to the port: %v", err)
	}
}
