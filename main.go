package main

import (
	"fmt"

	"github.com/pivovarit/hexagonal-go/app"
)

const appPort = 8080

func main() {
	mainApp := &app.App{}
	mainApp.Initialize()
	mainApp.Run(fmt.Sprintf(":%d", appPort))
}
