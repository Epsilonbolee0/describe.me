package handlers

import (
	ts "describe.me/internal/objects/transport"
	"describe.me/internal/service"
	"describe.me/internal/utils/messagez"
	. "describe.me/internal/utils/response"
	"describe.me/pkg/logger"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	Service *service.UserService
	Logger  logger.Interface
}

func NewUserHandler(rg *gin.RouterGroup, serv *service.UserService, log logger.Interface) {
	handler := &userHandler{serv, log}

	group := rg.Group("/auth")
	{
		group.GET("/login", handler.login)
		group.POST("/register", handler.register)
	}
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
func (handler *userHandler) login(c *gin.Context) {
	const logTrace = "http - v1 - login"
	var req ts.UserLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handler.Logger.Error(err, logTrace)
		WithError(c, err)
		return
	}

	user, err := handler.Service.Login(c.Request.Context(), &req)
	if err != nil {
		handler.Logger.Error(err, logTrace)
		WithError(c, err)
	}

	WithoutError(c, messagez.OK, ts.UserLoginResponse{User: user})
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
func (handler *userHandler) register(c *gin.Context) {
	const logTrace = "http - v1 - register"
	var req ts.UserRegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handler.Logger.Error(err, logTrace)
		WithError(c, err)
		return
	}

	err := handler.Service.Register(c.Request.Context(), &req)
	if err != nil {
		handler.Logger.Error(err, logTrace)
		WithError(c, err)
	}

	WithoutError(c, messagez.OK, ts.UserRegisterResponse{})
}
