package server

import (
	"fmt"
	"go-tsukamoto/internal/database"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"go-tsukamoto/internal/app/service/academic"
	"go-tsukamoto/internal/app/service/achievement"
	"go-tsukamoto/internal/app/service/activity"
	"go-tsukamoto/internal/app/service/course"
	fuzzy "go-tsukamoto/internal/app/service/fuzzy"
	"go-tsukamoto/internal/app/service/thesis"
	"go-tsukamoto/internal/app/service/user"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

type Server struct {
	port               int
	db                 database.Service
	userService        user.UserService
	achievementService achievement.AchievementService
	academicService    academic.AcademicService
	activityService    activity.ActivityService
	thesisService      thesis.ThesisService
	fuzzyService       fuzzy.FuzzyServiceInterface
	courseService      course.CourseServiceInterface
}

func NewServer(db *gorm.DB) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port:               port,
		db:                 database.New(),
		userService:        user.NewService(db),
		achievementService: achievement.NewService(db),
		academicService:    academic.NewService(db),
		activityService:    activity.NewService(db),
		thesisService:      thesis.NewService(db),
		fuzzyService:       fuzzy.NewService(db),
		courseService:      course.NewService(db),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.loggingMiddleware(NewServer.RegisterRoutes()),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Starting server on port %d\n", NewServer.port)
	return server
}

func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s %v\n", r.RemoteAddr, r.Method, r.URL, time.Since(start))
	})
}
