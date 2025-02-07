package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-processor-challenge/internal/models"
	"receipt-processor-challenge/internal/services"
	"receipt-processor-challenge/internal/storage"
	"strings"

	"github.com/google/uuid"
)

type ReceiptHandler struct {
	store   storage.Store
	service *services.PointsService
}

func NewReceiptHandler(store storage.Store, service *services.PointsService) *ReceiptHandler {
	return &ReceiptHandler{store: store, service: service}
}

func (h *ReceiptHandler) Process(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	var rc models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&rc); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	pts, err := h.service.Calculate(&rc)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	id := uuid.New().String()
	if err := h.store.Save(id, pts); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func (h *ReceiptHandler) Points(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/receipts/"), "/")
	if len(parts) != 2 || parts[1] != "points" {
		http.Error(w, "", http.StatusNotFound)
		return
	}
	pts, err := h.store.Get(parts[0])
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]int{"points": pts})
}

