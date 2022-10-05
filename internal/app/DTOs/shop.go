package DTOs

import "mime/multipart"

type CreateShop struct {
	Name          string                `form:"name" json:"name" validate:"required"`
	Logo          *multipart.FileHeader `form:"logo" json:"logo" validate:"omitempty"`
	LogoPath      string                `form:"-" json:"-"`
	Type          string                `form:"type" json:"type" validate:"required"`
	SocialAddress string                `form:"social_address" json:"social_address" validate:"required"`
	Description   string                `form:"description" json:"description" validate:"omitempty"`
	Phone         string                `form:"phone" json:"phone" validate:"omitempty"`
	Mobile        string                `form:"mobile" json:"mobile" validate:"omitempty,numeric,statswith=09"`
	TelegramID    string                `form:"telegram_id" json:"telegram_id" validate:"omitempty"`
	InstagramID   string                `form:"instagram_id" json:"instagram_id" validate:"omitempty"`
	WhatsappID    string                `form:"whatsapp_id" json:"whatsapp_id" validate:"omitempty"`
	Email         string                `form:"email" json:"email" validate:"omitempty,email"`
	Website       string                `form:"website" json:"website" validate:"omitempty"`
}

type UpdateShop struct {
	Name          string                `form:"name" json:"name" validate:"omitempty"`
	Logo          *multipart.FileHeader `form:"logo" json:"logo" validate:"omitempty"`
	LogoPath      string                `form:"-" json:"-"`
	LogoRemove    string                `form:"-" json:"-"`
	Type          string                `form:"type" json:"type" validate:"omitempty"`
	SocialAddress string                `form:"social_address" json:"social_address" validate:"omitempty"`
	Description   string                `form:"description" json:"description" validate:"omitempty"`
	Phone         string                `form:"phone" json:"phone" validate:"omitempty"`
	Mobile        string                `form:"mobile" json:"mobile" validate:"omitempty,numeric,statswith=09"`
	TelegramID    string                `form:"telegram_id" json:"telegram_id" validate:"omitempty"`
	InstagramID   string                `form:"instagram_id" json:"instagram_id" validate:"omitempty"`
	WhatsappID    string                `form:"whatsapp_id" json:"whatsapp_id" validate:"omitempty"`
	Email         string                `form:"email" json:"email" validate:"omitempty,email"`
	Website       string                `form:"website" json:"website" validate:"omitempty"`
	SendPrice     float32               `form:"send_price" json:"send_price" validate:"omitempty"`
	VerifySocial  bool                  `form:"-" json:"-" validate:"-"`
}

type CheckSocial struct {
	SocialAddress string `form:"social_address" json:"social_address" validate:"required"`
	Type          string `form:"type" json:"type" validate:"required"`
	ShopID        uint64 `form:"shop_id" json:"shop_id" validate:"required,numeric"`
}

type SendPrice struct {
	SendPrice float32 `form:"send_price" json:"send_price" validate:"required,numeric"`
	ShopID    uint64  `form:"shop_id" json:"shop_id" validate:"required,numeric"`
}
