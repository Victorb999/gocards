package main

import (
	"database/sql"
	"gocards/internal/service"
	"gocards/internal/web"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// go mod tidy
func main() {

	db, err := sql.Open("sqlite3", "cards.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	cardService := service.NewCardService(db)
	cardHandlers := web.NewCardHandlers(cardService)

	router := http.NewServeMux()

	router.HandleFunc("GET /cards", cardHandlers.GetCards)
	router.HandleFunc("GET /cards/{id}", cardHandlers.GetCardById)
	router.HandleFunc("POST /cards", cardHandlers.CreateCard)
	router.HandleFunc("DELETE /cards/{id}", cardHandlers.DeleteCard)
	router.HandleFunc("PUT /cards/{id}", cardHandlers.UpdateCard)

	http.ListenAndServe(":8080", router)
}
