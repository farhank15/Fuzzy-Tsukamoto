package handlers

import (
	"encoding/json"
	"errors"
	dto "go-tsukamoto/internal/app/dto/thesis"
	"go-tsukamoto/internal/app/service/thesis"
	"go-tsukamoto/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ThesisHandler struct {
	service thesis.ThesisService
}

func NewThesisHandler(service thesis.ThesisService) *ThesisHandler {
	return &ThesisHandler{service: service}
}

func (h *ThesisHandler) CreateThesis(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateThesisRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	resp, err := h.service.CreateThesis(r.Context(), &req)
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "Thesis created successfully", resp)
}

func (h *ThesisHandler) GetThesisByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid thesis ID", nil)
		return
	}
	resp, err := h.service.GetThesisByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFoundResponse(w, "Thesis not found")
		} else {
			utils.ServerErrorResponse(w, err)
		}
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "Thesis retrieved successfully", resp)
}

func (h *ThesisHandler) GetThesesByUserID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}
	resp, err := h.service.GetThesesByUserID(r.Context(), userID)
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "Theses retrieved successfully", resp)
}

func (h *ThesisHandler) GetAllTheses(w http.ResponseWriter, r *http.Request) {
	resp, err := h.service.GetAllTheses(r.Context())
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "All theses retrieved successfully", resp)
}

func (h *ThesisHandler) UpdateThesis(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateThesisRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid thesis ID", nil)
		return
	}
	resp, err := h.service.UpdateThesis(r.Context(), id, &req)
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "Thesis updated successfully", resp)
}

func (h *ThesisHandler) DeleteThesis(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid thesis ID", nil)
		return
	}
	if err := h.service.DeleteThesis(r.Context(), id); err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusNoContent, "Thesis deleted successfully", nil)
}
