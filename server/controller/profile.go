package controller

import (
	"encoding/json"
	"net/http"

	"../model/domain"
	"../model/service"
	"../utils"
	"github.com/gorilla/mux"
)

type ProfileController struct {
	profileService *service.ProfileService
}

func SetupProfileController(profileService *service.ProfileService, router *mux.Router) {
	controller := &ProfileController{profileService: profileService}
	router.HandleFunc("/profile/info", controller.Info).Methods("GET")
	router.HandleFunc("/profile/key", controller.CurrentKey).Methods("GET")

	router.HandleFunc("/profile/update/name", controller.UpdateName).Methods("PATCH")
	router.HandleFunc("/profile/update/group", controller.UpdateGroup).Methods("PATCH")
	router.HandleFunc("/profile/update/sex", controller.UpdateSex).Methods("PATCH")
}

func (controller *ProfileController) Info(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}

	if login, err := utils.LoginFromCookie(r); err != nil {
		resp = utils.NoCookie()
	} else {
		resp = controller.profileService.Info(login)
	}

	utils.Respond(w, resp)
}

func (controller *ProfileController) CurrentKey(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}

	if login, err := utils.LoginFromCookie(r); err != nil {
		resp = utils.NoCookie()
	} else {
		resp = controller.profileService.CurrentKey(login)
	}

	utils.Respond(w, resp)
}

func (controller *ProfileController) UpdateName(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.UserProfileDTO{}

	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		utils.Respond(w, utils.InvalidRequest())
	}

	if login, err := utils.LoginFromCookie(r); err != nil {
		resp = utils.NoCookie()
	} else {
		resp = controller.profileService.UpdateName(login, dto.Name)
	}

	utils.Respond(w, resp)
}

func (controller *ProfileController) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.UserProfileDTO{}

	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		utils.Respond(w, utils.InvalidRequest())
	}

	if login, err := utils.LoginFromCookie(r); err != nil {
		resp = utils.NoCookie()
	} else {
		resp = controller.profileService.UpdateGroup(login, dto.Group)
	}

	utils.Respond(w, resp)
}

func (controller *ProfileController) UpdateSex(w http.ResponseWriter, r *http.Request) {
	var resp map[string]interface{}
	dto := &domain.UserProfileDTO{}

	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		utils.Respond(w, utils.InvalidRequest())
	}

	if login, err := utils.LoginFromCookie(r); err != nil {
		resp = utils.NoCookie()
	} else {
		resp = controller.profileService.UpdateSex(login, dto.Sex)
	}

	utils.Respond(w, resp)
}
