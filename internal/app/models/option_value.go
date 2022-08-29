package models

type OptionValue struct {
	ID        int    `json:"id"`
	OptionID  int    `json:"option_id"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	Sort      uint   `json:"sort"`
	CreatedAt string `json:"created_at"`
}
