package app

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"pow-blockchain/config"
	"pow-blockchain/internal/controller/http"
	"pow-blockchain/internal/domain/service"
	"pow-blockchain/pkg/govalidator"
	"syscall"
)

func Run() {
	validator := govalidator.New()

	log := zerolog.New(os.Stdout)

	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	err = validator.Validate(context.Background(), cfg)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	log = log.Level(zerolog.Level(*cfg.Logger.Level)).With().Timestamp().Logger()

	defer log.Info().Msg("Application has been shut down")

	log.Debug().Msg("Loaded configuration")

	blockService := service.NewBlockService()
	blockController := http.NewBlockController(blockService, *validator)

	app := fiber.New()
	defer func() {
		app.Shutdown()
		log.Info().Msgf("Http server shut down successfully")
	}()
	apiGroup := app.Group("api")
	blockGroup := apiGroup.Group("block")

	blockController.RegisterRoutes(blockGroup)

	go func() {
		log.Info().Msgf("Start server on port: %s:%d", cfg.HttpServer.Host, cfg.HttpServer.Port)
		if err := app.Listen(fmt.Sprintf("%s:%d", cfg.HttpServer.Host, cfg.HttpServer.Port)); err != nil {
			log.Fatal().Msgf("Error starting Server: ", err)
		}
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit
}
