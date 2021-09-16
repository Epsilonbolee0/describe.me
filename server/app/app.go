package app

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"../controller"
	"../keygen"
	"../middleware"
	repo "../model/repository"
	"../model/service"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Run() {
	router := mux.NewRouter()
	router.Use(middleware.JwtAuthenithication)
	setupControllers(router)

	host := os.Getenv("server_host")
	port := os.Getenv("server_port")
	now := time.Now()

	fmt.Printf("[%s] Listening to %s:%s\n", now.Format("15:04:05"), host, port)

	go keygen.Generate()

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}

func setupControllers(router *mux.Router) {
	conn := GetConnection()

	setupAuthController(conn, router)
	setupProfileController(conn, router)
	setupFunctionController(conn, router)
	setupDescribeController(conn, router)
	setupLanguageController(conn, router)
}

func setupAuthController(conn *gorm.DB, router *mux.Router) {
	userRepo := repo.NewUserRepository(conn)
	authService := service.NewAuthService(userRepo)
	controller.SetupAuthController(authService, router)
}

func setupProfileController(conn *gorm.DB, router *mux.Router) {
	userRepo := repo.NewUserRepository(conn)
	profileService := service.NewProfileService(userRepo)
	controller.SetupProfileController(profileService, router)
}

func setupFunctionController(conn *gorm.DB, router *mux.Router) {
	functionRepo := repo.NewFunctionRepository(conn)
	functionService := service.NewFunctionService(functionRepo)
	controller.SetupFunctionController(functionService, router)
}

func setupDescribeController(conn *gorm.DB, router *mux.Router) {
	describeRepo := repo.NewDescriptionRepository(conn)
	describeService := service.NewDescribeService(describeRepo)
	controller.SetupDescribeController(describeService, router)
}

func setupLanguageController(conn *gorm.DB, router *mux.Router) {
	languageRepo := repo.NewLanguageRepository(conn)
	languageService := service.NewLanguageService(languageRepo)
	controller.SetupLanguageController(languageService, router)
}
