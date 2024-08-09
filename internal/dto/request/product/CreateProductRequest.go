package product

type CreateProductRequest struct {
	Name              string  `json:"name" validate:"required,min=3,max=12"`
	BaseCost          float64 `json:"baseCost" validate:"required"`
	ProductType       string  `json:"productType" validate:"required,min=3,max=12"`
	AdditionalKwhCost float64 `json:"additionalKwhCost" validate:"required"`
	IncludedKwh       float64 `json:"includedKwh" validate:"required"`
}
