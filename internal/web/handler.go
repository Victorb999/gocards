package web

import (
	"encoding/json"
	"gocards/internal/service"
	"net/http"
	"strconv"
)

type CardHandlers struct {
	service *service.CardService
}

func NewCardHandlers(service *service.CardService) *CardHandlers {
	return &CardHandlers{service: service}
}

func (h *CardHandlers) GetCards(w http.ResponseWriter, r *http.Request) {
	cards, err := h.service.GetCards()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cards)
}

func (h *CardHandlers) GetCardById(w http.ResponseWriter, r *http.Request) {
	cardStr := r.PathValue("id")
	id, err := strconv.Atoi(cardStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusInternalServerError)
		return
	}

	card, err := h.service.GetCard(id)
	if err != nil {
		http.Error(w, "não achou a carta", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(card)
}

func (h *CardHandlers) CreateCard(w http.ResponseWriter, r *http.Request) {
	var card service.Card
	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		http.Error(w, err.Error()+"deu ruim", http.StatusBadRequest)
		return
	}
	err = h.service.CreateCard(&card)
	if err != nil {
		http.Error(w, "deu ruim pra criar a carta", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(card)
}

func (h *CardHandlers) DeleteCard(w http.ResponseWriter, r *http.Request) {
	cardStr := r.PathValue("id")
	id, err := strconv.Atoi(cardStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusInternalServerError)
		return
	}
	err = h.service.DeleteCard(id)
	if err != nil {
		http.Error(w, "deu ruim pra apagar a carta", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *CardHandlers) UpdateCard(w http.ResponseWriter, r *http.Request) {
	cardStr := r.PathValue("id")
	id, err := strconv.Atoi(cardStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusInternalServerError)
		return
	}

	var card service.Card
	err = json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.UpdateCard(id, &card)
	if err != nil {
		http.Error(w, "deu ruim pra atualizar a carta", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(card)
	w.Header().Set("Content-Type", "application/json")
}
