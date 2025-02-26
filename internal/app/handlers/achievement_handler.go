package handlers

import (
	"encoding/json"
	"errors"
	dto "go-tsukamoto/internal/app/dto/achievement"
	"go-tsukamoto/internal/app/service/achievement"
	"go-tsukamoto/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AchievementHandler struct {
	service achievement.AchievementService
}

func NewAchievementHandler(service achievement.AchievementService) *AchievementHandler {
	return &AchievementHandler{service: service}
}

func (h *AchievementHandler) CreateAchievement(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateAchievementRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	resp, err := h.service.CreateAchievement(r.Context(), &req)
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "Achievement created successfully", resp)
}

func (h *AchievementHandler) GetAchievementByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid achievement ID", nil)
		return
	}
	resp, err := h.service.GetAchievementByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFoundResponse(w, "Achievement not found")
		} else {
			utils.ServerErrorResponse(w, err)
		}
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "Achievement retrieved successfully", resp)
}

func (h *AchievementHandler) GetAchievementsByUserID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}
	resp, err := h.service.GetAchievementsByUserID(r.Context(), userID)
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "Achievements retrieved successfully", resp)
}

func (h *AchievementHandler) UpdateAchievement(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateAchievementRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid achievement ID", nil)
		return
	}
	resp, err := h.service.UpdateAchievement(r.Context(), id, &req)
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "Achievement updated successfully", resp)
}

func (h *AchievementHandler) DeleteAchievement(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid achievement ID", nil)
		return
	}
	if err := h.service.DeleteAchievement(r.Context(), id); err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusNoContent, "Achievement deleted successfully", nil)
}
