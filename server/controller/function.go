package controller

import (
	"encoding/json"
	"net/http"

	"../model/domain"
	"../model/service"
	"../utils"
	"github.com/gorilla/mux"
)

type FunctionController struct {
	functionService *service.FunctionService
}

func SetupFunctionController(functionService *service.FunctionService, router *mux.Router) {
	controller := &FunctionController{functionService}

	router.HandleFunc("/function/add", controller.Create).Methods("POST")
	router.HandleFunc("/function/find", controller.Find).Methods("GET")
	router.HandleFunc("/function/rating", controller.Rating).Methods("GET")

	router.HandleFunc("/function/update/code", controller.UpdateCode).Methods("PATCH")
	router.HandleFunc("/function/like", controller.Like).Methods("PATCH")
	router.HandleFunc("/function/dislike", controller.Dislike).Methods("PATCH")
	router.HandleFunc("/function/indifferent", controller.Indifferent).Methods("PATCH")
}

func (controller *FunctionController) Create(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.FunctionDTO{}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		resp = utils.InvalidRequest()
	} else {
		resp = controller.functionService.Create(dto.LanguageID, dto.Code)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) Find(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.FunctionDTO{}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		resp = utils.InvalidRequest()
	} else {
		resp = controller.functionService.Find(dto.ID)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) UpdateCode(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.FunctionDTO{}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		resp = utils.InvalidRequest()
	} else {
		resp = controller.functionService.UpdateCode(dto.ID, dto.Code)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) Rating(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.FunctionDTO{}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		resp = utils.InvalidRequest()
	} else {
		resp = controller.functionService.Rating(dto.ID)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) Like(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.FunctionDTO{}

	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		utils.Respond(w, utils.InvalidRequest())
		return
	}

	if login, err := utils.LoginFromCookie(r); err != nil {
		resp = utils.NoCookie()
	} else {
		resp = controller.functionService.Like(login, dto.ID)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) Dislike(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.FunctionDTO{}

	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		utils.Respond(w, utils.InvalidRequest())
		return
	}

	if login, err := utils.LoginFromCookie(r); err != nil {
		resp = utils.NoCookie()
	} else {
		resp = controller.functionService.Dislike(login, dto.ID)
	}

	utils.Respond(w, resp)
}

func (controller *FunctionController) Indifferent(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.FunctionDTO{}

	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		utils.Respond(w, utils.InvalidRequest())
		return
	}

	if login, err := utils.LoginFromCookie(r); err != nil {
		resp = utils.NoCookie()
	} else {
		resp = controller.functionService.Indifferent(login, dto.ID)
	}

	utils.Respond(w, resp)
}
