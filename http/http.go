package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nwuw/dev/http/api"
	m "github.com/nwuw/dev/http/middleware"
	"go.uber.org/zap"
)

func Init() {
	r := gin.Default()
	r.Use(m.Cors())
	api.IndexApi(r)
	
	if err := r.Run(fmt.Sprintf(":%d", 924)); err != nil {
		zap.L().Sugar().Errorf("server run failed, err:%v\n", err)
		return
	}
}
