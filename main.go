package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

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

	values, err := url.ParseQuery(string(body))
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	bot, err := tgbotapi.NewBotAPI("7653357171:AAH0irLiLcA8TvZTXHNY3f_pCZMzVPMN0_E")
	chatID := int64(414747434) // ID бота
	//chatID := int64(-1002561300903) // ID тестового чата
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	count := 0
	for _, value := range values {
		m := value[count]

		msg := tgbotapi.NewMessage(chatID, m)
		bot.Send(msg)
		count++
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(values)
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
