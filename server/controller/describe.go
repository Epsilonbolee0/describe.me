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

	router.HandleFunc("/description/add", controller.Create).Methods("POST")
	router.HandleFunc("/description/list", controller.ListByFunction).Methods("GET")
	router.HandleFunc("/description/rating", controller.Rating).Methods("GET")
	router.HandleFunc("/description/update/code", controller.Delete).Methods("DELETE")

	router.HandleFunc("/description/like", controller.Like).Methods("PATCH")
	router.HandleFunc("/description/dislike", controller.Dislike).Methods("PATCH")
	router.HandleFunc("/description/indifferent", controller.Indifferent).Methods("PATCH")
}

func (controller *DescribeController) Create(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.DescribeDTO{}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		resp = utils.InvalidRequest()
	} else {
		resp = controller.describeService.Create(dto.ID, dto.Content)
	}

	utils.Respond(w, resp)
}

func (controller *DescribeController) ListByFunction(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.DescribeDTO{}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		resp = utils.InvalidRequest()
	} else {
		resp = controller.describeService.ListByFunction(dto.ID)
	}

	utils.Respond(w, resp)
}

func (controller *DescribeController) Delete(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.DescribeDTO{}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		resp = utils.InvalidRequest()
	} else {
		resp = controller.describeService.Delete(dto.ID)
	}

	utils.Respond(w, resp)
}

func (controller *DescribeController) Rating(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.DescribeDTO{}

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		resp = utils.InvalidRequest()
	} else {
		resp = controller.describeService.Rating(dto.ID)
	}

	utils.Respond(w, resp)
}

func (controller *DescribeController) Like(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.DescribeDTO{}

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
	dto := &domain.DescribeDTO{}

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
	dto := &domain.DescribeDTO{}

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
