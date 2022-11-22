package DTOs

type UpdateUser struct {
	//شناسه کاربر برای ویرایش
	ID uint64 `json:"id" form:"id" binding:"required" example:"1"`
	//جنسیت
	Gender string `form:"gender" json:"gender" validate:"omitempty,oneof=man woman" example:"man" enums:"man,woman"`
	//نام
	Firstname string `form:"firstname" json:"firstname" validate:"omitempty" example:"امیر"`
	//نام خانوادگی
	Lastname string `form:"lastname" json:"lastname" validate:"omitempty" example:"شیردل"`
	//ایمیل
	Email string `form:"email" json:"email" validate:"omitempty,email" example:"amirex128@gmail.com"`
	//موبایل
	Mobile string `form:"mobile" json:"mobile" validate:"omitempty,startswith=09,numeric" example:"09024809750"`
	//شماره کارت
	CartNumber string `form:"cart_number" json:"cart_number" validate:"omitempty,numeric,min=16,max=16" example:"6037998125410760"`
	//شماره شبا
	Shaba string `form:"shaba" json:"shaba" validate:"omitempty,numeric,min=24,max=24" example:"IR820540102680020817909002"`
	//پسورد
	Password string `form:"password" json:"password"  validate:"omitempty,min=6,max=20" example:"123456789"`
	//تکرار پسورد در صورت پر شدن به قصد تغییر پسورد
	AgainPassword string `form:"again_password" json:"again_password"  validate:"omitempty,min=6,max=20,eqfield=Password" example:"123456789"`
}
