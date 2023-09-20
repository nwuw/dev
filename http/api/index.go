package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nwuw/dev/http/router"
)

func Index(r *gin.Engine) {
	index := r.Group("/")
	
	index.GET("", router.Index)
}
