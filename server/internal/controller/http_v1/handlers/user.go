package handlers

import (
	ts "describe.me/internal/objects/transport"
	"describe.me/internal/service"
	"describe.me/internal/utils/messagez"
	. "describe.me/internal/utils/response"
	"describe.me/pkg/logger"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type userHandler struct {
	Service *service.UserService
	Logger  logger.Interface
}

func NewUserHandler(router *chi.Mux, serv *service.UserService, log logger.Interface) {
	handler := &userHandler{serv, log}

	router.Route("/auth", func(r chi.Router) {
		r.Get("/login", handler.login)
		r.Post("/register", handler.register)
	})
}

// @Summary     Login
// @Description Login as student
// @ID          login
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Param       request body ts.UserLoginRequest true "comment"
// @Success     200 {object} ts.UserLoginResponse
// @Failure     403 {object} resp
// @Failure 	500 {object} resp
// @Router      /auth/login [get]
func (handler *userHandler) login(w http.ResponseWriter, r *http.Request) {
	const logTrace = "http - v1 - login"
	var req ts.UserLoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.Logger.Error(err, logTrace)
		WithError(w, err)
		return
	}

	user, err := handler.Service.Login(r.Context(), &req)
	if err != nil {
		handler.Logger.Error(err, logTrace)
		WithError(w, err)
	}

	WithoutError(w, messagez.OK, ts.UserLoginResponse{User: user})
}

// @Summary     Register
// @Description Register a student
// @ID          register
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Param       request body ts.UserRegisterRequest true "comment"
// @Success     200 {object} ts.UserRegisterResponse
// @Failure     403 {object} resp
// @Failure     500 {object} resp
// @Router      /auth/register [post]
func (handler *userHandler) register(w http.ResponseWriter, r *http.Request) {
	const logTrace = "http - v1 - register"
	var req ts.UserRegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.Logger.Error(err, logTrace)
		WithError(w, err)
		return
	}

	err := handler.Service.Register(r.Context(), &req)
	if err != nil {
		handler.Logger.Error(err, logTrace)
		WithError(w, err)
	}

	WithoutError(w, messagez.OK, ts.UserRegisterResponse{})
}
