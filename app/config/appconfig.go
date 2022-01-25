package config

import (
	"context"
	"fmt"
	"net/http"
	"project2/app/internal"

	"github.com/gin-gonic/gin"
)

type AppConfiguration struct {
	config   Config
	engine   *gin.Engine
	server   *http.Server
	project2 internal.Project2Service
}

func NewAppService(config Config) *AppConfiguration {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	app := &AppConfiguration{
		config: config,
		engine: engine,
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", config.Port),
			Handler: engine,
		},
	}
	initializeServices(app)
	return app
}

func (a *AppConfiguration) Init() (err error) {
	a.initialiseRoutes()
	return nil
}

func (a *AppConfiguration) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		a.server.Shutdown(ctx)
	}()

	if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
