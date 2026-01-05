package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"

	"my-go-project/internal/config"
	"my-go-project/internal/db"
	"my-go-project/internal/handler"
	"my-go-project/internal/repository"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	gormDB, err := db.NewDB(cfg)
	if err != nil {
		log.Fatalf("connect db: %v", err)
	}

	repo := repository.NewMasterCategoryMerchantRepository(gormDB)

	app := fiber.New()
	handler.RegisterMasterCategoryMerchantRoutes(app, repo)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Listen(cfg.Addr); err != nil {
			log.Printf("server stopped: %v", err)
			shutdown <- syscall.SIGTERM
		}
	}()

	<-shutdown
	if err := app.Shutdown(); err != nil {
		log.Printf("shutdown error: %v", err)
	}
}
