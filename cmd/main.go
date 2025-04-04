package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"municipality_app/cmd/container"
	"time"
)

func initTime() {
	time.Local = time.UTC
}

func runServer(r *gin.Engine) {
	go func() {
		if err := r.Run(":8080"); err != nil {
			panic(err)
		}
	}()
}
func main() {
	initTime()

	fx.New(
		container.MainContainer,
		fx.Invoke(runServer),
	).Run()
}
