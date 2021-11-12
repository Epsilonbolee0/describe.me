package response

import "github.com/gin-gonic/gin"

type resp struct {
	Error string `json:"error" example:"message"`
}

func WithError(c *gin.Context, err error) {
	msg := err.Error()
	c.AbortWithStatusJSON(StatusByMessage(msg), resp{msg})
}

func WithoutError(c *gin.Context, msg string, obj interface{}) {
	c.JSON(StatusByMessage(msg), obj)
}
