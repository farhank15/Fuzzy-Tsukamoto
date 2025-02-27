package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-tsukamoto/internal/app/dto/academic"
	"go-tsukamoto/internal/app/handlers"
	mockAcademicService "go-tsukamoto/internal/app/service/academic"
	"go-tsukamoto/utils"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateAcademic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mockAcademicService.NewMockAcademicService(ctrl)
	handler := handlers.NewAcademicHandler(mockService)

	t.Run("Success", func(t *testing.T) {
		reqBody := &academic.CreateAcademicRequest{
			UserID:          1,
			Ipk:             3.5,
			RepeatedCourses: 0,
			Semester:        1,
			Year:            2023,
			PredicateID:     1,
		}
		reqBodyBytes, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/academic", bytes.NewReader(reqBodyBytes))
		w := httptest.NewRecorder()

		mockService.EXPECT().CreateAcademic(gomock.Any(), reqBody).Return(&academic.AcademicResponse{
			ID:              1,
			UserID:          reqBody.UserID,
			Ipk:             reqBody.Ipk,
			RepeatedCourses: reqBody.RepeatedCourses,
			Semester:        reqBody.Semester,
			Year:            reqBody.Year,
			PredicateID:     &reqBody.PredicateID,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}, nil)

		handler.CreateAcademic(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var resp utils.ResponseData
		json.NewDecoder(w.Body).Decode(&resp)
		assert.Equal(t, "Academic record created successfully", resp.Message)
	})

	t.Run("Invalid Request Body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/academic", bytes.NewReader([]byte("invalid body")))
		w := httptest.NewRecorder()

		handler.CreateAcademic(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var resp utils.ResponseData
		json.NewDecoder(w.Body).Decode(&resp)
		assert.Equal(t, "invalid character 'i' looking for beginning of value", resp.Message)
	})

	t.Run("Service Error", func(t *testing.T) {
		reqBody := &academic.CreateAcademicRequest{
			UserID:          1,
			Ipk:             3.5,
			RepeatedCourses: 0,
			Semester:        1,
			Year:            2023,
			PredicateID:     1,
		}
		reqBodyBytes, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/academic", bytes.NewReader(reqBodyBytes))
		w := httptest.NewRecorder()

		mockService.EXPECT().CreateAcademic(gomock.Any(), reqBody).Return(nil, errors.New("service error"))

		handler.CreateAcademic(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var resp utils.ResponseData
		json.NewDecoder(w.Body).Decode(&resp)
		assert.Equal(t, "Internal Server Error", resp.Message)
	})
}

func TestGetAcademicByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mockAcademicService.NewMockAcademicService(ctrl)
	handler := handlers.NewAcademicHandler(mockService)

	t.Run("Success", func(t *testing.T) {
		academicID := 1
		req := httptest.NewRequest(http.MethodGet, "/academic/"+strconv.Itoa(academicID), nil)
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(academicID)})

		mockService.EXPECT().GetAcademicByID(gomock.Any(), academicID).Return(&academic.AcademicResponse{
			ID:              academicID,
			UserID:          1,
			Ipk:             3.5,
			RepeatedCourses: 0,
			Semester:        1,
			Year:            2023,
			PredicateID:     nil,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}, nil)

		handler.GetAcademicByID(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var resp utils.ResponseData
		json.NewDecoder(w.Body).Decode(&resp)
		assert.Equal(t, "Academic record retrieved successfully", resp.Message)
	})

	t.Run("Invalid Academic ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/academic/invalid", nil)
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": "invalid"})

		handler.GetAcademicByID(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var resp utils.ResponseData
		json.NewDecoder(w.Body).Decode(&resp)
		assert.Equal(t, "Invalid academic ID", resp.Message)
	})

	t.Run("Academic Not Found", func(t *testing.T) {
		academicID := 999
		req := httptest.NewRequest(http.MethodGet, "/academic/"+strconv.Itoa(academicID), nil)
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(academicID)})

		mockService.EXPECT().GetAcademicByID(gomock.Any(), academicID).Return(nil, gorm.ErrRecordNotFound)

		handler.GetAcademicByID(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		var resp utils.ResponseData
		json.NewDecoder(w.Body).Decode(&resp)
		assert.Equal(t, "Academic record not found", resp.Message)
	})

	t.Run("Service Error", func(t *testing.T) {
		academicID := 1
		req := httptest.NewRequest(http.MethodGet, "/academic/"+strconv.Itoa(academicID), nil)
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(academicID)})

		mockService.EXPECT().GetAcademicByID(gomock.Any(), academicID).Return(nil, errors.New("service error"))

		handler.GetAcademicByID(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var resp utils.ResponseData
		json.NewDecoder(w.Body).Decode(&resp)
		assert.Equal(t, "Internal Server Error", resp.Message)
	})
}

func TestGetAcademicsByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mockAcademicService.NewMockAcademicService(ctrl)
	handler := handlers.NewAcademicHandler(mockService)

	t.Run("Success", func(t *testing.T) {
		userID := 1
		req := httptest.NewRequest(http.MethodGet, "/academic/user/"+strconv.Itoa(userID), nil)
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"user_id": strconv.Itoa(userID)})

		mockService.EXPECT().GetAcademicsByUserID(gomock.Any(), userID).Return([]*academic.AcademicResponse{
			{
				ID:              1,
				UserID:          userID,
				Ipk:             3.5,
				RepeatedCourses: 0,
				Semester:        1,
				Year:            2023,
				PredicateID:     nil,
				CreatedAt:       time.Now(),
				UpdatedAt:       time.Now(),
			},
			{
				ID:              2,
				UserID:          userID,
				Ipk:             3.6,
				RepeatedCourses: 0,
				Semester:        2,
				Year:            2023,
				PredicateID:     nil,
				CreatedAt:       time.Now(),
				UpdatedAt:       time.Now(),
			},
		}, nil)

		handler.GetAcademicsByUserID(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var resp utils.ResponseData
		json.NewDecoder(w.Body).Decode(&resp)
		assert.Equal(t, "Academic records retrieved successfully", resp.Message)
	})

	t.Run("Invalid User ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/academic/user/invalid", nil)
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"user_id": "invalid"})

		handler.GetAcademicsByUserID(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var resp utils.ResponseData
		json.NewDecoder(w.Body).Decode(&resp)
		assert.Equal(t, "Invalid user ID", resp.Message)
	})

	t.Run("Service Error", func(t *testing.T) {
		userID := 1
		req := httptest.NewRequest(http.MethodGet, "/academic/user/"+strconv.Itoa(userID), nil)
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"user_id": strconv.Itoa(userID)})

		mockService.EXPECT().GetAcademicsByUserID(gomock.Any(), userID).Return(nil, errors.New("service error"))

		handler.GetAcademicsByUserID(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var resp utils.ResponseData
		json.NewDecoder(w.Body).Decode(&resp)
		assert.Equal(t, "Internal Server Error", resp.Message)
	})
}

func TestGetAllAcademics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mockAcademicService.NewMockAcademicService(ctrl)
	handler := handlers.NewAcademicHandler(mockService)

	t.Run("Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/academic", nil)
		w := httptest.NewRecorder()

		mockService.EXPECT().GetAllAcademics(gomock.Any()).Return([]*academic.AcademicResponse{
			{
				ID:              1,
				UserID:          1,
				Ipk:             3.5,
				RepeatedCourses: 0,
				Semester:        1,
				Year:            2023,
				PredicateID:     nil,
				CreatedAt:       time.Now(),
				UpdatedAt:       time.Now(),
			},
			{
				ID:              2,
				UserID:          2,
				Ipk:             3.6,
				RepeatedCourses: 0,
				Semester:        2,
				Year:            2023,
				PredicateID:     nil,
				CreatedAt:       time.Now(),
				UpdatedAt:       time.Now(),
			},
		}, nil)

		handler.GetAllAcademics(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var resp utils.ResponseData
		json.NewDecoder(w.Body).Decode(&resp)
		assert.Equal(t, "All academic records retrieved successfully", resp.Message)
	})

	t.Run("Service Error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/academic", nil)
		w := httptest.NewRecorder()

		mockService.EXPECT().GetAllAcademics(gomock.Any()).Return(nil, errors.New("service error"))

		handler.GetAllAcademics(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var resp utils.ResponseData
		json.NewDecoder(w.Body).Decode(&resp)
		assert.Equal(t, "Internal Server Error", resp.Message)
	})
}

func TestUpdateAcademic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mockAcademicService.NewMockAcademicService(ctrl)
	handler := handlers.NewAcademicHandler(mockService)

	t.Run("Success", func(t *testing.T) {
		reqBody := &academic.UpdateAcademicRequest{
			Ipk:             3.75,
			RepeatedCourses: 1,
			Semester:        2,
			Year:            2023,
			PredicateID:     2,
		}
		reqBodyBytes, _ := json.Marshal(reqBody)
		academicID := 1
		req := httptest.NewRequest(http.MethodPut, "/academic/"+strconv.Itoa(academicID), bytes.NewReader(reqBodyBytes))
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(academicID)})

		mockService.EXPECT().UpdateAcademic(gomock.Any(), academicID, reqBody).Return(&academic.AcademicResponse{
			ID:              academicID,
			UserID:          1,
			Ipk:             reqBody.Ipk,
			RepeatedCourses: reqBody.RepeatedCourses,
			Semester:        reqBody.Semester,
			Year:            reqBody.Year,
			PredicateID:     &reqBody.PredicateID,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}, nil)

		handler.UpdateAcademic(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var resp utils.ResponseData
		json.NewDecoder(w.Body).Decode(&resp)
		assert.Equal(t, "Academic record updated successfully", resp.Message)
	})

	t.Run("Invalid Request Body", func(t *testing.T) {
		academicID := 1
		req := httptest.NewRequest(http.MethodPut, "/academic/"+strconv.Itoa(academicID), bytes.NewReader([]byte("invalid body")))
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(academicID)})

		handler.UpdateAcademic(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var resp utils.ResponseData
		json.NewDecoder(w.Body).Decode(&resp)
		assert.Equal(t, "invalid character 'i' looking for beginning of value", resp.Message)
	})

	t.Run("Invalid Academic ID", func(t *testing.T) {
		reqBody := &academic.UpdateAcademicRequest{
			Ipk:             3.75,
			RepeatedCourses: 1,
			Semester:        2,
			Year:            2023,
			PredicateID:     2,
		}
		reqBodyBytes, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPut, "/academic/invalid", bytes.NewReader(reqBodyBytes))
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": "invalid"})

		handler.UpdateAcademic(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var resp utils.ResponseData
		json.NewDecoder(w.Body).Decode(&resp)
		assert.Equal(t, "Invalid academic ID", resp.Message)
	})

	t.Run("Service Error", func(t *testing.T) {
		reqBody := &academic.UpdateAcademicRequest{
			Ipk:             3.75,
			RepeatedCourses: 1,
			Semester:        2,
			Year:            2023,
			PredicateID:     2,
		}
		reqBodyBytes, _ := json.Marshal(reqBody)
		academicID := 1
		req := httptest.NewRequest(http.MethodPut, "/academic/"+strconv.Itoa(academicID), bytes.NewReader(reqBodyBytes))
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(academicID)})

		mockService.EXPECT().UpdateAcademic(gomock.Any(), academicID, reqBody).Return(nil, errors.New("service error"))

		handler.UpdateAcademic(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var resp utils.ResponseData
		json.NewDecoder(w.Body).Decode(&resp)
		assert.Equal(t, "Internal Server Error", resp.Message)
	})
}

func TestDeleteAcademic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mockAcademicService.NewMockAcademicService(ctrl)
	handler := handlers.NewAcademicHandler(mockService)

	t.Run("Success", func(t *testing.T) {
		academicID := 1
		req := httptest.NewRequest(http.MethodDelete, "/academic/"+strconv.Itoa(academicID), nil)
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(academicID)})

		mockService.EXPECT().DeleteAcademic(gomock.Any(), academicID).Return(nil)

		handler.DeleteAcademic(w, req)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("Invalid Academic ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/academic/invalid", nil)
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": "invalid"})

		handler.DeleteAcademic(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var resp utils.ResponseData
		json.NewDecoder(w.Body).Decode(&resp)
		assert.Equal(t, "Invalid academic ID", resp.Message)
	})

	t.Run("Service Error", func(t *testing.T) {
		academicID := 1
		req := httptest.NewRequest(http.MethodDelete, "/academic/"+strconv.Itoa(academicID), nil)
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(academicID)})

		mockService.EXPECT().DeleteAcademic(gomock.Any(), academicID).Return(errors.New("service error"))

		handler.DeleteAcademic(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var resp utils.ResponseData
		json.NewDecoder(w.Body).Decode(&resp)
		assert.Equal(t, "Internal Server Error", resp.Message)
	})
}
