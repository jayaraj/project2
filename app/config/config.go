package config

import (
	"net/http"
	"project2/app/internal/service"

	"github.com/jayaraj/project1/pkg"
)

type Config struct {
	Port        int
	Project1Url string
}

func initializeServices(appConfig *AppConfiguration) {
	project1 := pkg.NewProject1(appConfig.config.Project1Url, http.DefaultClient)
	appConfig.project2 = service.NewProject2Service(project1)
}
