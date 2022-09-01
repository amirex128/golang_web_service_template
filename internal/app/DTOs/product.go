package DTOs

type IndexProduct struct {
	Search   string `form:"search" json:"search" validate:"min=3,max=50"`
	Page     int    `form:"page" json:"page"`
	PageSize int    `form:"page_size" json:"page_size"`
	Sort     string `form:"sort" json:"sort"`
}
type CreateProduct struct {
	ManufacturerId   int     `json:"manufacturer_id"`
	Description      string  `json:"description"`
	Name             string  `json:"name"`
	ShortDescription string  `json:"short_description"`
	Quantity         int     `json:"quantity"`
	Price            float32 `json:"price"`
	FreeSend         string  `json:"free_send"`
	Weight           uint    `json:"weight"`
	Height           uint    `json:"height"`
	Width            uint    `json:"width"`
	Active           byte    `json:"active"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
	StartedAt        string  `json:"started_at"`
	EndedAt          string  `json:"ended_at"`
	DeliveryTime     string  `json:"delivery_time"` // مدت زمان ارسال
	OptionId         int     `json:"option_id"`
	OptionItemID     int     `json:"option_item_id"`
}
