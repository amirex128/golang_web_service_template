package models

type Guild struct {
	ID         int    `json:"id"`
	ParentID   int    `json:"parent_id"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	Equivalent string `json:"equivalent"`
	Sort       uint   `json:"sort"`
	Active     byte   `json:"active"`
}

type GuildProduct struct {
	GuildID   int   `json:"guild_id"`
	ProductID int64 `json:"product_id"`
}
