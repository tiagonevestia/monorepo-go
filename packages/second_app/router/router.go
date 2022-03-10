package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tiagonevestia/monorepo-go/packages/shared/handlers/health"
)

func GetEngine() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/", health.Health("Second app"))
	return r
}
