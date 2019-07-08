package app

import (
	"ginDemo/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(http.StatusOK, gin.H{
		"code": httpCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})
	return
}
