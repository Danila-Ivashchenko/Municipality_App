package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"municipality_app/cmd/container"
	"municipality_app/internal/common/migrator"
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

func upMigrations(m *migrator.BaseMigrator) error {
	return m.Up()
}

func main() {
	initTime()

	fx.New(
		container.MainContainer,
		fx.Invoke(upMigrations),
		fx.Invoke(runServer),
	).Run()
}
