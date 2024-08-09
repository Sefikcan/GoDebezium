package main

import (
	"database/sql"
	"github.com/sefikcan/godebezium/internal/server"
	"github.com/sefikcan/godebezium/pkg/config"
	"github.com/sefikcan/godebezium/pkg/logger"
	"github.com/sefikcan/godebezium/pkg/storage/postgres"
)

func main() {
	cfg := config.NewConfig()

	zapLogger := logger.NewLogger(cfg)
	zapLogger.InitLogger()
	zapLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode, false)

	psqlDB, err := postgres.NewPsqlDB(cfg)
	db, err := psqlDB.DB()
	if err != nil {
		zapLogger.Fatalf("Postgresql init: %s", err)
	} else {
		zapLogger.Infof("Postgres connected, Status: %#v", db.Stats())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			zapLogger.Fatalf("Postgresql init: %s", err)
		}
	}(db)

	s := server.NewServer(cfg, psqlDB, zapLogger)
	if err = s.Run(); err != nil {
		zapLogger.Fatal(err)
	}
}
