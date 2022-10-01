package models

type Financial struct {
	ID     uint64 `json:"id"`
	UserID uint64 `json:"user_id"`
	User   User   `gorm:"foreignKey:user_id" json:"user"`
	Amount uint64 `json:"amount"`
}
