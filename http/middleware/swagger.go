package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nwuw/dev/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerInit(r *gin.Engine) {

	// set swagger info
	docs.SwaggerInfo.Title = "dev API"
	docs.SwaggerInfo.Description = "https://github.com/xxandjg"
	docs.SwaggerInfo.Version = "v0.0.1"
	docs.SwaggerInfo.Host = "localhost:924"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
