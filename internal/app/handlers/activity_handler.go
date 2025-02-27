package handlers

import (
	"encoding/json"
	"errors"
	dto "go-tsukamoto/internal/app/dto/activity"
	"go-tsukamoto/internal/app/service/activity"
	"go-tsukamoto/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ActivityHandler struct {
	service activity.ActivityService
}

func NewActivityHandler(service activity.ActivityService) *ActivityHandler {
	return &ActivityHandler{service: service}
}

func (h *ActivityHandler) CreateActivity(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateActivityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	resp, err := h.service.CreateActivity(r.Context(), &req)
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "Activity created successfully", resp)
}

func (h *ActivityHandler) GetActivityByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid activity ID", nil)
		return
	}
	resp, err := h.service.GetActivityByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFoundResponse(w, "Activity not found")
		} else {
			utils.ServerErrorResponse(w, err)
		}
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "Activity retrieved successfully", resp)
}

func (h *ActivityHandler) GetActivitiesByUserID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}
	resp, err := h.service.GetActivitiesByUserID(r.Context(), userID)
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "Activities retrieved successfully", resp)
}

func (h *ActivityHandler) GetAllActivities(w http.ResponseWriter, r *http.Request) {
	resp, err := h.service.GetAllActivities(r.Context())
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "All activities retrieved successfully", resp)
}

func (h *ActivityHandler) UpdateActivity(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateActivityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid activity ID", nil)
		return
	}
	resp, err := h.service.UpdateActivity(r.Context(), id, &req)
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "Activity updated successfully", resp)
}

func (h *ActivityHandler) DeleteActivity(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid activity ID", nil)
		return
	}
	if err := h.service.DeleteActivity(r.Context(), id); err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusNoContent, "Activity deleted successfully", nil)
}
