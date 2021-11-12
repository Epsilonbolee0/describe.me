package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"../model/domain"
	"../model/service"
	"../utils"
)

type AuthController struct {
	authService *service.AuthService
}

func SetupAuthController(authService *service.AuthService, router *mux.Router) {
	controller := &AuthController{authService: authService}
	router.HandleFunc("/auth/login", controller.Login).Methods("POST")
	router.HandleFunc("/auth/logout", controller.Logout).Methods("POST")
	router.HandleFunc("/auth/register", controller.Register).Methods("POST")
}

func (controller *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.UserAuthDTO{}

	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		resp = utils.InvalidRequest()
	} else {
		resp = controller.authService.Login(dto.Login, dto.Password)
		utils.SetTokenCookie(w, dto.Login)
	}

	utils.Respond(w, resp)
}

func (controller *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	utils.DiscardTokenCookie(w)
	utils.Respond(w, utils.Message("Logout was successful"))
}

func (controller *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.UserAuthDTO{}

	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		resp = utils.InvalidRequest()
	} else {
		resp = controller.authService.Register(dto.Key, dto.Login, dto.Password)
		utils.SetTokenCookie(w, dto.Login)
	}

	utils.Respond(w, resp)
}
