package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func (a *App) Initialize() {
	a.router = gin.Default()
	a.router.GET("/health", func(context *gin.Context) {
		context.Header("Content-Type", "application/json")
		_, _ = context.Writer.Write([]byte(`{"status": "UP"}`))
	})
}

func (a *App) Run(port string) {
	err := a.router.Run(port)
	if err != nil {
		panic(fmt.Sprintf("Couldn't start application %v", err))
	}
}
