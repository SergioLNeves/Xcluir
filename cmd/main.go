package main

import (
	"github.com/SergioLNeves/Xcluir/config"
	"github.com/SergioLNeves/Xcluir/handler"
	"github.com/SergioLNeves/Xcluir/repository"
	"github.com/SergioLNeves/Xcluir/services"
	"github.com/gookit/slog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	configureLogger()

	configureTweetRoutes(e)

	log.Println("Servidor iniciado na porta 8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

func configureLogger() {
	slog.Configure(func(logger *slog.SugaredLogger) {
		f := logger.Formatter.(*slog.TextFormatter)
		f.EnableColor = true
	})
}

func configureTweetRoutes(e *echo.Echo) {
	tweetRepo := repository.NewTweetRepository()
	tweetService := services.NewTweetServices(tweetRepo)
	tweetHandler := handler.NewTweetHandler(tweetService)

	v1 := e.Group("/v1")
	v1.POST("/delete-tweets/upload", tweetHandler.DeleteTweetsFromFile)
	v1.DELETE("/delete-tweets/:filepath", tweetHandler.DeleteTweetsFromFile)
}
