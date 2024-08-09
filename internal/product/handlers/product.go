package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sefikcan/godebezium/internal/dto/request/product"
	"github.com/sefikcan/godebezium/internal/product/usecase"
	"github.com/sefikcan/godebezium/pkg/config"
	"github.com/sefikcan/godebezium/pkg/logger"
	"github.com/sefikcan/godebezium/pkg/util"
	"net/http"
	"strings"
)

type ProductHandlers interface {
	Create() echo.HandlerFunc
}

type productHandlers struct {
	cfg            *config.Config
	logger         logger.Logger
	productUseCase usecase.ProductUseCase
}

func (p productHandlers) Create() echo.HandlerFunc {
	return func(e echo.Context) error {
		productRequest := product.CreateProductRequest{}
		if err := e.Bind(&productRequest); err != nil {
			return e.JSON(http.StatusBadRequest, util.NewHttpResponse(http.StatusBadRequest, strings.ToLower(err.Error()), nil))
		}

		createdProduct, err := p.productUseCase.Create(context.Background(), productRequest)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, util.NewHttpResponse(http.StatusInternalServerError, strings.ToLower(err.Error()), nil))
		}

		return e.JSON(http.StatusCreated, createdProduct)
	}
}

func NewProductHandlers(cfg *config.Config, logger logger.Logger, productUseCase usecase.ProductUseCase) ProductHandlers {
	return &productHandlers{
		cfg:            cfg,
		logger:         logger,
		productUseCase: productUseCase,
	}
}
