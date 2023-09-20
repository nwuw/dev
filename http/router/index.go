package router

import (
	"github.com/gin-gonic/gin"
	k "github.com/nwuw/dev/kit"
	s "github.com/nwuw/dev/service"
)

var is = s.IndexService{}

// Index godoc
// @Summary      index
// @Description  系统首页后端路由
// @Tags         index
// @Produce      json
// @Success      200  {object}  response.IndexResponse
// @Failure      200  {object}  response.IndexResponse
// @Router       / [get]
func Index(c *gin.Context) {
	k.Ok(c, is.Index())
}
