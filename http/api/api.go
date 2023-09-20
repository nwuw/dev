package api

import (
	"github.com/gin-gonic/gin"
	m "github.com/nwuw/dev/http/middleware"
)

func IndexApi(e *gin.Engine) {
	m.Init(e)
	Index(e)
}
