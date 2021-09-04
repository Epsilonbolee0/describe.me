package controller

import (
	"encoding/json"
	"net/http"

	"../model/domain"
	"../model/service"
	"../utils"
	"github.com/gorilla/mux"
)

type LanguageController struct {
	languageService *service.LanguageService
}

func SetupLanguageController(languageService *service.LanguageService, router *mux.Router) {
	controller := &LanguageController{languageService}

	router.HandleFunc("/lang/list", controller.List).Methods("GET")
	router.HandleFunc("/lang/find", controller.Find).Methods("GET")
	router.HandleFunc("/lang/add", controller.Create).Methods("POST")
	router.HandleFunc("/lang/delete", controller.Delete).Methods("DELETE")
}

func (controller *LanguageController) List(w http.ResponseWriter, r *http.Request) {
	utils.Respond(w, controller.languageService.List())
}

func (controller *LanguageController) Find(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.LanguageDTO{}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		resp = utils.InvalidRequest()
	} else {
		resp = controller.languageService.Find(dto.Name)
	}

	utils.Respond(w, resp)
}

func (controller *LanguageController) Create(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.LanguageDTO{}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		resp = utils.InvalidRequest()
	} else {
		resp = controller.languageService.Create(dto.Name)
	}

	utils.Respond(w, resp)
}

func (controller *LanguageController) Delete(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.LanguageDTO{}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		resp = utils.InvalidRequest()
	} else {
		resp = controller.languageService.Delete(dto.ID)
	}

	utils.Respond(w, resp)
}
