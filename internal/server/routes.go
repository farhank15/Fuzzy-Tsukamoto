package server

import (
	"encoding/json"
	"go-tsukamoto/internal/app/handlers"
	"go-tsukamoto/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) RegisterRoutes() http.Handler {
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/health", s.healthHandler).Methods("GET")

	// User routes
	userHandler := handlers.NewUserHandler(s.userService)
	router.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
	router.HandleFunc("/login", userHandler.LoginUser).Methods("POST")
	router.HandleFunc("/user/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", userHandler.GetUserByID).Methods("GET")
	router.HandleFunc("/user/{id}", userHandler.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/detail/{id}", userHandler.GetUserWithRelatedData).Methods("GET")

	// Achievement routes
	achievementHandler := handlers.NewAchievementHandler(s.achievementService)
	router.HandleFunc("/achievement", achievementHandler.CreateAchievement).Methods("POST")
	router.HandleFunc("/achievement/{id}", achievementHandler.GetAchievementByID).Methods("GET")
	router.HandleFunc("/achievement/user/{user_id}", achievementHandler.GetAchievementsByUserID).Methods("GET")
	router.HandleFunc("/achievement", achievementHandler.GetAllAchievements).Methods("GET")
	router.HandleFunc("/achievement/{id}", achievementHandler.UpdateAchievement).Methods("PUT")
	router.HandleFunc("/achievement/{id}", achievementHandler.DeleteAchievement).Methods("DELETE")

	// Academic routes
	academicHandler := handlers.NewAcademicHandler(s.academicService)
	router.HandleFunc("/academic", academicHandler.CreateAcademic).Methods("POST")
	router.HandleFunc("/academic/{id}", academicHandler.GetAcademicByID).Methods("GET")
	router.HandleFunc("/academic/user/{user_id}", academicHandler.GetAcademicsByUserID).Methods("GET")
	router.HandleFunc("/academic", academicHandler.GetAllAcademics).Methods("GET")
	router.HandleFunc("/academic/{id}", academicHandler.UpdateAcademic).Methods("PUT")
	router.HandleFunc("/academic/{id}", academicHandler.DeleteAcademic).Methods("DELETE")

	// Activity routes
	activityHandler := handlers.NewActivityHandler(s.activityService)
	router.HandleFunc("/activity", activityHandler.CreateActivity).Methods("POST")
	router.HandleFunc("/activity/{id}", activityHandler.GetActivityByID).Methods("GET")
	router.HandleFunc("/activity/user/{user_id}", activityHandler.GetActivitiesByUserID).Methods("GET")
	router.HandleFunc("/activity", activityHandler.GetAllActivities).Methods("GET")
	router.HandleFunc("/activity/{id}", activityHandler.UpdateActivity).Methods("PUT")
	router.HandleFunc("/activity/{id}", activityHandler.DeleteActivity).Methods("DELETE")

	// Thesis routes
	thesisHandler := handlers.NewThesisHandler(s.thesisService)
	router.HandleFunc("/thesis", thesisHandler.CreateThesis).Methods("POST")
	router.HandleFunc("/thesis/{id}", thesisHandler.GetThesisByID).Methods("GET")
	router.HandleFunc("/thesis/student/{student_id}", thesisHandler.GetThesesByUserID).Methods("GET")
	router.HandleFunc("/thesis", thesisHandler.GetAllTheses).Methods("GET")
	router.HandleFunc("/thesis/{id}", thesisHandler.UpdateThesis).Methods("PUT")
	router.HandleFunc("/thesis/{id}", thesisHandler.DeleteThesis).Methods("DELETE")

	// Fuzzy route
	fuzzyHandler := handlers.NewFuzzyHandler(s.fuzzyService)
	router.HandleFunc("/fuzzy", fuzzyHandler.CalculateFuzzy).Methods("POST")

	// Course routes
	courseHandler := handlers.NewCourseHandler(s.courseService)
	router.HandleFunc("/course", courseHandler.CreateCourse).Methods("POST")
	router.HandleFunc("/courses/import", courseHandler.ImportCourses).Methods("POST")
	router.HandleFunc("/course/{id}", courseHandler.GetCourseByID).Methods("GET")
	router.HandleFunc("/course", courseHandler.GetCourses).Methods("GET")
	router.HandleFunc("/course/{id}", courseHandler.UpdateCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", courseHandler.DeleteCourse).Methods("DELETE")

	// Wrap the router with CORS middleware
	return middleware.CorsMiddleware(router)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(s.db.Health())
	if err != nil {
		http.Error(w, "Failed to marshal health check response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
