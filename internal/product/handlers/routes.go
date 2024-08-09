package handlers

import "github.com/labstack/echo/v4"

func NewProductRoutes(productRouteGroup *echo.Group, p ProductHandlers) {
	productRouteGroup.POST("", p.Create())
}
