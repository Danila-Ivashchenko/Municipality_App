package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"municipality_app/cmd/container"
	"municipality_app/internal/common/config"
	"municipality_app/internal/common/migrator"
	"time"
)

func initTime() {
	time.Local = time.UTC
}

func runServer(r *gin.Engine, cfg *config.Config) {
	if !cfg.UseHttp() && !cfg.UseHttps() {
		panic("server dont`t listen http and https")
	}

	if cfg.UseHttp() {
		go func() {
			if err := r.Run(fmt.Sprintf("%s:%s", cfg.HttpHost(), cfg.HttpPort())); err != nil {
				panic(err)
			}
		}()
	}

	if cfg.UseHttps() {
		go func() {
			if err := r.RunTLS(fmt.Sprintf("%s:%s", cfg.HttpsHost(), cfg.HttpsPort()), cfg.CertPath(), cfg.CertKeyPath()); err != nil {
				panic(err)
			}
		}()
	}
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
