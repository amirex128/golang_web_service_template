package models

type FeatureItem struct {
	ID             int    `json:"id"`
	FeatureGroupID int    `json:"feature_group_id"`
	Title          string `json:"title"`
	Type           string `json:"type"`    // label color boolean other  => تعیین نوع مجاز پاسخ برای این خصوصیت : پاسخ متنی ساده - پاسخ از نوع رنگ - پاسخ بله یا خیر - پاسخ چند جوابه
	Actions        string `json:"actions"` // search top_feature short_desc point => خاص تر کردن یک خصوصیت و تعیین اینکه این خصوصیت عملکرد و استفاده برای جستجو دارد و خصوصیت برجسته است یا خصوصیت برای استفاده در توضیحات مختصر است و یا خصوصیت برای نظردهی است
	Active         byte   `json:"active"`
	Icon           string `json:"icon"`
	InputType      string `json:"input_type"` // number text float all => نوع ورودی های مجاز برای مقدار ویژگی ها
	Sort           uint   `json:"sort"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
