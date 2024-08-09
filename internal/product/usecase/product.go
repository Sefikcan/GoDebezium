package usecase

import (
	"context"
	"github.com/pkg/errors"
	request "github.com/sefikcan/godebezium/internal/dto/request/product"
	response "github.com/sefikcan/godebezium/internal/dto/response/product"
	"github.com/sefikcan/godebezium/internal/product/mapping"
	"github.com/sefikcan/godebezium/internal/product/repository"
	"github.com/sefikcan/godebezium/pkg/config"
	"github.com/sefikcan/godebezium/pkg/logger"
	"github.com/sefikcan/godebezium/pkg/util"
	"net/http"
)

type ProductUseCase interface {
	Create(ctx context.Context, request request.CreateProductRequest) (*response.ProductResponse, error)
}

type productUseCase struct {
	cfg               *config.Config
	productRepository repository.ProductRepository
	logger            logger.Logger
}

func (p productUseCase) Create(ctx context.Context, request request.CreateProductRequest) (*response.ProductResponse, error) {
	if err := util.ValidateStruct(&request); err != nil {
		return nil, util.NewHttpResponse(http.StatusBadRequest, util.BadRequest.Error(), errors.WithMessage(err, "productUseCase.Create.ValidateStruct"))
	}

	product := mapping.CreateDtoToEntity(&request)

	resp, err := p.productRepository.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	mappedResponse := mapping.MapDto(resp)

	return mappedResponse, nil
}

func NewProductUseCase(cfg *config.Config, productRepository repository.ProductRepository, logger logger.Logger) ProductUseCase {
	return &productUseCase{
		cfg:               cfg,
		productRepository: productRepository,
		logger:            logger,
	}
}
