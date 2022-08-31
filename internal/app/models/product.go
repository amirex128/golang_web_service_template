package models

import (
	"encoding/gob"
	"io"
)

type Product struct {
	ID               int64   `json:"id"`
	UserId           string  `json:"user_id"`
	ManufacturerId   int     `json:"manufacturer_id"`
	GuildId          int     `json:"guild_id"`
	Description      string  `json:"description"`
	Name             string  `json:"name"`
	Barcode          string  `json:"barcode"`
	ShortDescription string  `json:"short_description"`
	Unit             string  `json:"unit"` // adad halghe m2 m cm day hour minute second kg g mesghal litre pors nafar year month week visit
	MadeInIran       byte    `json:"made_in_iran"`
	Consumer         string  `json:"consumer"` // man woman little_girl little_boy baby_girl baby_boy old_man old_woman company use_in_home use_in_public_place plant animal
	TotalSales       string  `json:"total_sales"`
	TotalVisitors    string  `json:"total_visitors"`
	ProductType      string  `json:"product_type"` // normal virtual nobatdehi ads sayer zarfiatvamohlat roozmoshakhas
	BlockStatus      string  `json:"block_status"` // ok filter black_list disable_ok disable_filter disable_black_list
	OptionId         int     `json:"option_id"`
	CityID           int     `json:"city_id"`
	ProvinceID       int     `json:"province_id"`
	MapAreaID        int     `json:"map_area_id"`
	Quantity         int     `json:"quantity"`
	MinCount         int     `json:"min_count"`
	MaxCount         int     `json:"max_count"`
	Price            float32 `json:"price"`
	WrapPrice        float32 `json:"wrap_price"` // هزینه دسته بندی
	FreeSend         string  `json:"free_send"`
	Tax              float32 `json:"tax"`
	Porsant          float32 `json:"porsant"`
	PorsantValue     float32 `json:"porsant_value"`
	Weight           uint    `json:"weight"`
	Height           uint    `json:"height"`
	Wide             uint    `json:"wide"`
	Width            uint    `json:"width"`
	Condition        string  `json:"condition"` // new used recycled new_hurt
	Rank             string  `json:"rank"`
	HasGuarantee     string  `json:"has_guarantee"`
	Active           string  `json:"active"`
	Lat              float64 `json:"lat"`
	Lon              float64 `json:"lon"`
	TypeBuy          string  `json:"type_buy"` // online cod both نوع پرداخت سفارش آنلاین یا پرداخت در محل یا هر دو
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
	StartedAt        string  `json:"started_at"`
	EndedAt          string  `json:"ended_at"`
	DeletedAt        string  `json:"deleted_at"`
	Delivery         string  `json:"delivery"` // مدت زمان ارسال
	IsOriginal       string  `json:"is_original"`
	OptionItemID     int     `json:"option_item_id"`
}
type ProductArr []Product

func (s ProductArr) Len() int {
	return len(s)
}
func (s ProductArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ProductArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *Product) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *Product) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
