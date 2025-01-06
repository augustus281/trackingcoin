package dto

type ListingParam struct {
	Start int `json:"start" validate:"gte=1"`
	Limit int `json:"limi" validate:"min=1,max=5000"`
}
