package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`    // code infor for user
	Message string      `json:"message"` // message infor for user
	Data    interface{} `json:"data"`    // dynamic data
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    ErrorCodeSuccess,
		Message: msg[ErrorCodeSuccess],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
