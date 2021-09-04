package controller

import (
	"encoding/json"
	"net/http"

	"../model/domain"
	"../model/service"
	"../utils"
	"github.com/gorilla/mux"
)

type DescribeController struct {
	describeService *service.DescribeService
}

func SetupDescribeController(describeService *service.DescribeService, router *mux.Router) {
	controller := &DescribeController{describeService}

	router.HandleFunc("/describe/add", controller.Create).Methods("POST")
	router.HandleFunc("/describe/find", controller.Find).Methods("GET")
	router.HandleFunc("/describe/rating", controller.Rating).Methods("GET")

	router.HandleFunc("/describe/update/code", controller.UpdateCode).Methods("PATCH")
	router.HandleFunc("/describe/like", controller.Like).Methods("PATCH")
	router.HandleFunc("/describe/dislike", controller.Dislike).Methods("PATCH")
	router.HandleFunc("/describe/indifferent", controller.Indifferent).Methods("PATCH")
}

func (controller *DescribeController) Create(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.FunctionDescribeDTO{}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		resp = utils.InvalidRequest()
	} else {
		resp = controller.describeService.Create(dto.LanguageID, dto.Code)
	}

	utils.Respond(w, resp)
}

func (controller *DescribeController) Find(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.FunctionDescribeDTO{}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		resp = utils.InvalidRequest()
	} else {
		resp = controller.describeService.Find(dto.ID)
	}

	utils.Respond(w, resp)
}

func (controller *DescribeController) UpdateCode(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.FunctionDescribeDTO{}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		resp = utils.InvalidRequest()
	} else {
		resp = controller.describeService.UpdateCode(dto.ID, dto.Code)
	}

	utils.Respond(w, resp)
}

func (controller *DescribeController) Rating(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.FunctionDescribeDTO{}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		resp = utils.InvalidRequest()
	} else {
		resp = controller.describeService.Rating(dto.ID)
	}

	utils.Respond(w, resp)
}

func (controller *DescribeController) Like(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.FunctionDescribeDTO{}

	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		utils.Respond(w, utils.InvalidRequest())
		return
	}

	if login, err := utils.LoginFromCookie(r); err != nil {
		resp = utils.NoCookie()
	} else {
		resp = controller.describeService.Like(login, dto.ID)
	}

	utils.Respond(w, resp)
}

func (controller *DescribeController) Dislike(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.FunctionDescribeDTO{}

	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		utils.Respond(w, utils.InvalidRequest())
		return
	}

	if login, err := utils.LoginFromCookie(r); err != nil {
		resp = utils.NoCookie()
	} else {
		resp = controller.describeService.Dislike(login, dto.ID)
	}

	utils.Respond(w, resp)
}

func (controller *DescribeController) Indifferent(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.FunctionDescribeDTO{}

	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		utils.Respond(w, utils.InvalidRequest())
		return
	}

	if login, err := utils.LoginFromCookie(r); err != nil {
		resp = utils.NoCookie()
	} else {
		resp = controller.describeService.Indifferent(login, dto.ID)
	}

	utils.Respond(w, resp)
}
