package handlers

import (
	"encoding/json"
	"errors"
	dto "go-tsukamoto/internal/app/dto/academic"
	"go-tsukamoto/internal/app/service/academic"
	"go-tsukamoto/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AcademicHandler struct {
	service academic.AcademicService
}

func NewAcademicHandler(service academic.AcademicService) *AcademicHandler {
	return &AcademicHandler{service: service}
}

func (h *AcademicHandler) CreateAcademic(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateAcademicRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	resp, err := h.service.CreateAcademic(r.Context(), &req)
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "Academic record created successfully", resp)
}

func (h *AcademicHandler) GetAcademicByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid academic ID", nil)
		return
	}
	resp, err := h.service.GetAcademicByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFoundResponse(w, "Academic record not found")
		} else {
			utils.ServerErrorResponse(w, err)
		}
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "Academic record retrieved successfully", resp)
}

func (h *AcademicHandler) GetAcademicsByUserID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}
	resp, err := h.service.GetAcademicsByUserID(r.Context(), userID)
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "Academic records retrieved successfully", resp)
}

func (h *AcademicHandler) GetAllAcademics(w http.ResponseWriter, r *http.Request) {
	resp, err := h.service.GetAllAcademics(r.Context())
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "All academic records retrieved successfully", resp)
}

func (h *AcademicHandler) UpdateAcademic(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateAcademicRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid academic ID", nil)
		return
	}
	resp, err := h.service.UpdateAcademic(r.Context(), id, &req)
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "Academic record updated successfully", resp)
}

func (h *AcademicHandler) DeleteAcademic(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid academic ID", nil)
		return
	}
	if err := h.service.DeleteAcademic(r.Context(), id); err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusNoContent, "Academic record deleted successfully", nil)
}
