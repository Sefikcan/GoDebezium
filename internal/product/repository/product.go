package repository

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sefikcan/godebezium/internal/product/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(ctx context.Context, product entity.Product) (entity.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func (p productRepository) Create(ctx context.Context, product entity.Product) (entity.Product, error) {
	if result := p.db.WithContext(ctx).Create(&product); result.Error != nil {
		return entity.Product{}, errors.Wrap(result.Error, "productRepository.Create.DbError")
	}

	return product, nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}
