package mapping

import (
	"github.com/sefikcan/godebezium/internal/dto/response/product"
	"github.com/sefikcan/godebezium/internal/product/entity"
)

func MapDto(p entity.Product) *product.ProductResponse {
	return &product.ProductResponse{
		Id:                p.Id,
		Name:              p.Name,
		ProductType:       p.ProductType,
		BaseCost:          p.BaseCost,
		AdditionalKwhCost: p.AdditionalKwhCost,
		IncludedKwh:       p.IncludedKwh,
	}
}
