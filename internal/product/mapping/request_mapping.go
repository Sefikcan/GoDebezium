package mapping

import (
	"github.com/sefikcan/godebezium/internal/dto/request/product"
	"github.com/sefikcan/godebezium/internal/product/entity"
)

func CreateDtoToEntity(p *product.CreateProductRequest) entity.Product {
	return entity.Product{
		Name:              p.Name,
		ProductType:       p.ProductType,
		BaseCost:          p.BaseCost,
		AdditionalKwhCost: p.AdditionalKwhCost,
		IncludedKwh:       p.IncludedKwh,
	}
}
