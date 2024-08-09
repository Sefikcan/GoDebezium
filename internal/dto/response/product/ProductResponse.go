package product

type ProductResponse struct {
	Id                int     `json:"id"`
	Name              string  `json:"name"`
	BaseCost          float64 `json:"baseCost"`
	ProductType       string  `json:"productType"`
	AdditionalKwhCost float64 `json:"additionalKwhCost"`
	IncludedKwh       float64 `json:"includedKwh"`
}
