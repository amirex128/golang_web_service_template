package models

type Category struct {
	ID          int    `json:"id"`
	ParentID    int    `json:"parent_id"`
	Name        string `json:"name"`
	Sort        uint   `json:"sort"`
	Equivalent  string `json:"equivalent"`
	GuildIds    string `json:"guild_ids"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Depth       byte   `json:"depth"`
	Parents     string `json:"parents"`
	Active      byte   `json:"active"`
}

type CategoryOption struct {
	CategoryID int `json:"category_id"`
	OptionID   int `json:"option_id"`
}

type CategoryProduct struct {
	CategoryID int   `json:"category_id"`
	ProductID  int64 `json:"product_id"`
}

type CategoryRelated struct {
	CategoryID        int `json:"category_id"`
	CategoryRelatedID int `json:"category_related_id"`
}
