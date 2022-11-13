package DTOs

type CreateShop struct {
	Name          string  `form:"name" json:"name" validate:"required"`
	EnglishName   string  `form:"english_name" json:"english_name" validate:"required"`
	GalleryID     uint64  `form:"gallery_id" json:"gallery_id" validate:"omitempty"`
	Type          string  `form:"type" json:"type" validate:"required,oneof=instagram telegram website rubika"`
	SocialAddress string  `form:"social_address" json:"social_address" validate:"required"`
	Description   string  `form:"description" json:"description" validate:"omitempty"`
	Phone         string  `form:"phone" json:"phone" validate:"omitempty"`
	Mobile        string  `form:"mobile" json:"mobile" validate:"omitempty,numeric,startswith=09"`
	TelegramID    string  `form:"telegram_id" json:"telegram_id" validate:"omitempty"`
	InstagramID   string  `form:"instagram_id" json:"instagram_id" validate:"omitempty"`
	WhatsappID    string  `form:"whatsapp_id" json:"whatsapp_id" validate:"omitempty"`
	Email         string  `form:"email" json:"email" validate:"omitempty,email"`
	Website       string  `form:"website" json:"website" validate:"omitempty"`
	SendPrice     float32 `form:"send_price" json:"send_price" validate:"omitempty"`
	ThemeID       uint64  `form:"theme_id" json:"theme_id" validate:"omitempty"`
}

type UpdateShop struct {
	Name          string  `form:"name" json:"name" validate:"omitempty"`
	EnglishName   string  `form:"english_name" json:"english_name" validate:"omitempty"`
	GalleryID     uint64  `form:"gallery_id" json:"gallery_id" validate:"omitempty"`
	Type          string  `form:"type" json:"type" validate:"omitempty,oneof=instagram telegram website rubika"`
	SocialAddress string  `form:"social_address" json:"social_address" validate:"omitempty"`
	Description   string  `form:"description" json:"description" validate:"omitempty"`
	Phone         string  `form:"phone" json:"phone" validate:"omitempty"`
	Mobile        string  `form:"mobile" json:"mobile" validate:"omitempty,numeric,startswith=09"`
	TelegramID    string  `form:"telegram_id" json:"telegram_id" validate:"omitempty"`
	InstagramID   string  `form:"instagram_id" json:"instagram_id" validate:"omitempty"`
	WhatsappID    string  `form:"whatsapp_id" json:"whatsapp_id" validate:"omitempty"`
	Email         string  `form:"email" json:"email" validate:"omitempty,email"`
	Website       string  `form:"website" json:"website" validate:"omitempty"`
	SendPrice     float32 `form:"send_price" json:"send_price" validate:"omitempty"`
	VerifySocial  bool    `form:"-" json:"-" validate:"-"`
}

type DeleteShop struct {
	NewShopID     uint64 `form:"new_shop_id" json:"new_shop_id" validate:"omitempty"`
	ProductBehave string `form:"product_behave" json:"product_behave" validate:"oneof=move delete_product,required"`
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

type IndexShop struct {
	Index
}
