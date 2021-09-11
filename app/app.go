package app

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pivovarit/hexagonal-go/app/account"
	"github.com/pivovarit/hexagonal-go/app/config"
)

type App struct {
	router   *gin.Engine
	config   *config.AppConfiguration
	accounts *account.Service
}

func (a *App) Initialize() {
	a.accounts = &account.Service{}
	a.router = gin.Default()
	a.router.GET("/health", func(context *gin.Context) {
		context.Header("Content-Type", "application/json")
		response, _ := json.Marshal(map[string]interface{}{"status": "UP"})
		_, _ = context.Writer.Write(response)
	})
}

func (a *App) Run(port string) {
	err := a.router.Run(port)
	if err != nil {
		panic(fmt.Sprintf("Couldn't start application %v", err))
	}
}
