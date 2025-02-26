package handlers

import (
	"encoding/json"
	"errors"
	dto "go-tsukamoto/internal/app/dto/course"
	"go-tsukamoto/internal/app/service/course"
	"go-tsukamoto/utils"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type CourseHandler struct {
	service course.CourseServiceInterface
}

func NewCourseHandler(service course.CourseServiceInterface) *CourseHandler {
	return &CourseHandler{service: service}
}

func (h *CourseHandler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCourseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	resp, err := h.service.CreateCourse(r.Context(), &req)
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "Course created successfully", resp)
}

func (h *CourseHandler) GetCourseByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid course ID", nil)
		return
	}
	resp, err := h.service.GetCourseByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || err.Error() == "course not found" {
			utils.NotFoundResponse(w, "Course not found")
		} else {
			utils.ServerErrorResponse(w, err)
		}
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "Course retrieved successfully", resp)
}

func (h *CourseHandler) GetCourses(w http.ResponseWriter, r *http.Request) {
	resp, err := h.service.GetCourses(r.Context())
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "Courses retrieved successfully", resp)
}

func (h *CourseHandler) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateCourseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid course ID", nil)
		return
	}
	resp, err := h.service.UpdateCourse(r.Context(), id, &req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || err.Error() == "course not found" {
			utils.NotFoundResponse(w, "Course not found")
		} else {
			utils.ServerErrorResponse(w, err)
		}
		return
	}
	utils.SuccessResponse(w, http.StatusOK, "Course updated successfully", resp)
}

func (h *CourseHandler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid course ID", nil)
		return
	}
	if err := h.service.DeleteCourse(r.Context(), id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || err.Error() == "course not found" {
			utils.NotFoundResponse(w, "Course not found")
		} else {
			utils.ServerErrorResponse(w, err)
		}
		return
	}
	utils.SuccessResponse(w, http.StatusNoContent, "Course deleted successfully", nil)
}

func (h *CourseHandler) ImportCourses(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PathFile string `json:"path_file"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	file, err := os.Open(req.PathFile)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Failed to open file", nil)
		return
	}
	defer file.Close()

	var courses []dto.CreateCourseRequest
	if err := json.NewDecoder(file).Decode(&courses); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Failed to decode JSON file", nil)
		return
	}

	for _, req := range courses {
		if _, err := h.service.CreateCourse(r.Context(), &req); err != nil {
			utils.ServerErrorResponse(w, err)
			return
		}
	}

	utils.SuccessResponse(w, http.StatusOK, "Courses imported successfully", nil)
}
