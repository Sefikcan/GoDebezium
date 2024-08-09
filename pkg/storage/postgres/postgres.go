package postgres

import (
	"fmt"
	"github.com/sefikcan/godebezium/internal/product/entity"
	"github.com/sefikcan/godebezium/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func NewPsqlDB(c *config.Config) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		c.Postgres.UserName,
		c.Postgres.Password,
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.DbName)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Otomatik tablo oluşturma
	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(c.Postgres.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(c.Postgres.ConnMaxLifeTime) * time.Second)
	sqlDB.SetMaxIdleConns(c.Postgres.MaxIdleConns)
	sqlDB.SetConnMaxIdleTime(time.Duration(c.Postgres.ConnMaxIdleTime) * time.Second)
	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
