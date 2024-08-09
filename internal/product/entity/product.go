package entity

type Product struct {
	Id                int     `gorm:"primary_key" json:"id"`
	Name              string  `gorm:"index:idx_name,unique" json:"name"`
	BaseCost          float64 `json:"baseCost"`
	ProductType       string  `json:"productType"`
	AdditionalKwhCost float64 `json:"additionalKwhCost"`
	IncludedKwh       float64 `json:"includedKwh"`
}
