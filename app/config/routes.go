package config

import (
	"project2/app/internal/httpservice"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (a *AppConfiguration) initialiseRoutes() {
	a.engine.Use(gin.Recovery())
	a.engine.Use(cors.Default())
	v1 := a.engine.Group("api/v1")
	a.addV1Routes(v1)
}

func (a *AppConfiguration) addV1Routes(router *gin.RouterGroup) {
	router.POST("/toptenwords", httpservice.CreateTopTenUsedWordsHandler(a.project2))
}
