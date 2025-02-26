package handlers

import (
	"encoding/json"

	dto "go-tsukamoto/internal/app/dto/fuzzy"
	service "go-tsukamoto/internal/app/service/fuzzy"
	"go-tsukamoto/utils"
	"net/http"
)

type FuzzyHandler struct {
	service service.FuzzyServiceInterface
}

func NewFuzzyHandler(service service.FuzzyServiceInterface) *FuzzyHandler {
	return &FuzzyHandler{service: service}
}

func (h *FuzzyHandler) CalculateFuzzy(w http.ResponseWriter, r *http.Request) {
	var req dto.FuzzyRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	resp, err := h.service.CalculateFuzzy(r.Context(), req.UserID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Fuzzy calculation successful", resp)
}
