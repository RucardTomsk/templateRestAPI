package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
	"templateRestAPI/cmd/api/controller"
	"templateRestAPI/cmd/api/router"
	"templateRestAPI/cmd/config"
	"templateRestAPI/cmd/storage/migration"
	"templateRestAPI/internal/common"
	"templateRestAPI/internal/server"
	"templateRestAPI/internal/telemetry/log"
)

func main() {
	logger := log.NewLogger()

	appCli := common.InitAppCli()
	if err := appCli.Run(os.Args); err != nil {
		logger.Fatal(err.Error())
	}

	// read config
	var cfg config.Config
	if err := viper.MergeInConfig(); err != nil {
		logger.Fatal(fmt.Sprintf("error reading config file: %v", err))
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		logger.Fatal(fmt.Sprintf("unable to decode into struct: %v", err))
	}

	// init connections
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Fatal(fmt.Sprintf("can't connect to database: %v", err))
	}

	logger.Info(fmt.Sprintf("successfully connected to database %s on %s:%d as %s",
		cfg.DB.Name, cfg.DB.Host, cfg.DB.Port, cfg.DB.User))

	if err := migration.Migrate(db); err != nil {
		logger.Fatal(fmt.Sprintf("failed to migrate database: %v", err))
	}
	logger.Info("database migrated successfully")

	// init storage

	// init services

	// init controllers
	controllers := controller.NewControllerContainer(
		logger)

	// init server

	handler := router.NewRouter(cfg)
	srv := new(server.Server)

	go func() {
		if err := srv.Run(cfg.Server.Host, cfg.Server.Port, handler.InitRoutes(
			logger,
			controllers,
		)); err != nil {
			logger.Error(fmt.Sprintf("error accured while running http server: %s", err.Error()))
		}
	}()

	logger.Info(fmt.Sprintf("listening on %s:%s", cfg.Server.Host, cfg.Server.Port))

	// handle signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("shutting down gracefully...")
	defer func() { logger.Info("shutdown complete") }()

	// perform shutdown
	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Error(fmt.Sprintf("error occured on server shutting down: %s", err.Error()))
	}
}
