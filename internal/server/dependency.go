package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sefikcan/godebezium/internal/product/handlers"
	"github.com/sefikcan/godebezium/internal/product/repository"
	"github.com/sefikcan/godebezium/internal/product/usecase"
	"net/http"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	productRepository := repository.NewProductRepository(s.db)

	productUseCase := usecase.NewProductUseCase(s.cfg, productRepository, s.logger)

	productHandler := handlers.NewProductHandlers(s.cfg, s.logger, productUseCase)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID},
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10, //1kb
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("2M"))

	v1 := e.Group("/api/v1")
	health := v1.Group("/health")
	productGroup := v1.Group("/products")

	handlers.NewProductRoutes(productGroup, productHandler)

	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	return nil
}
