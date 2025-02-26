package handlers

import (
	"encoding/json"
	"errors"
	dto "go-tsukamoto/internal/app/dto/user"
	"go-tsukamoto/internal/app/service/user"
	"go-tsukamoto/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type UserHandler struct {
	service user.UserService
}

func NewUserHandler(service user.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	resp, err := h.service.CreateUser(r.Context(), &req)
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "User registered successfully", resp)
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	user, err := h.service.GetUserByUsername(r.Context(), req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.ErrorResponse(w, http.StatusUnauthorized, "Invalid username or password", nil)
		} else {
			utils.ServerErrorResponse(w, err)
		}
		return
	}
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		utils.ErrorResponse(w, http.StatusUnauthorized, "Invalid username or password", nil)
		return
	}
	token, err := utils.GenerateJWT(user.ID, user.Name, user.Nim)
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "Login successful", map[string]string{"access": token})
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if h.service == nil {
		utils.ServerErrorResponse(w, errors.New("user service is not initialized"))
		return
	}
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	resp, err := h.service.CreateUser(r.Context(), &req)
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	resp.Password = ""
	utils.SuccessResponse(w, http.StatusOK, "User created successfully", resp)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}
	resp, err := h.service.UpdateUser(r.Context(), id, &req)
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	resp.Password = "" // Exclude password from response
	utils.SuccessResponse(w, http.StatusOK, "User updated successfully", resp)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}
	resp, err := h.service.GetUserByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFoundResponse(w, "User not found")
		} else {
			utils.ServerErrorResponse(w, err)
		}
		return
	}
	// Include hashed password in response for debugging
	utils.SuccessResponse(w, http.StatusOK, "User retrieved successfully", resp)
}

func (h *UserHandler) GetUserWithRelatedData(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}
	resp, err := h.service.GetUserWithRelatedData(r.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFoundResponse(w, "User not found")
		} else {
			utils.ServerErrorResponse(w, err)
		}
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "User retrieved successfully", resp)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}
	if err := h.service.DeleteUser(r.Context(), id); err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusNoContent, "User deleted successfully", nil)
}
