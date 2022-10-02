package DTOs

import "mime/multipart"

type CreateShop struct {
	Name       string `form:"name" json:"name" validate:"required"`
	Type       string `form:"type" json:"type" validate:"required,oneof=instagram telegram website"`
	Social     string `form:"social" json:"social" validate:"required"`
	CategoryID uint64 `form:"category_id" json:"category_id" validate:"required,numeric"`
}

type UpdateShop struct {
	Name        string                `form:"name" json:"name" validate:"required"`
	Logo        *multipart.FileHeader `form:"logo" json:"logo" validate:"omitempty"`
	CategoryID  uint64                `form:"category_id" json:"category_id" validate:"omitempty,numeric"`
	Description string                `form:"description" json:"description" validate:"omitempty"`
	Phone       string                `form:"phone" json:"phone" validate:"omitempty"`
	Mobile      string                `form:"mobile" json:"mobile" validate:"omitempty,startswith=09"`
	TelegramID  string                `form:"telegram_id" json:"telegram_id" validate:"omitempty"`
	InstagramID string                `form:"instagram_id" json:"instagram_id" validate:"omitempty"`
	WhatsappID  string                `form:"whatsapp_id" json:"whatsapp_id" validate:"omitempty"`
	Email       string                `form:"email" json:"email" validate:"omitempty"`
	Website     string                `form:"website" json:"website" validate:"omitempty"`
}
