package DTOs

type Verify struct {
	//موبایل
	Mobile string `form:"mobile" json:"mobile" validate:"required,min=11,max=11,startswith=09" example:"09024809750"`
	//کد تایید درصورتی که با کد میخواهید وارد شوید
	VerifyCode string `form:"verify_code" json:"verify_code" validate:"omitempty,min=4,max=4" example:"4563"`
	//پسورد درصورتی که با پسورد میخواهید وارد شودی
	Password string `form:"password" json:"password"  validate:"omitempty,min=6,max=20" example:"123456789"`
}
type RequestLoginRegister struct {
	//موبایل
	Mobile string `form:"mobile" json:"mobile"  validate:"required,min=11,max=11,startswith=09" example:"09024809750"`
}

type ChangePassword struct {
	//پسورد
	Password string `form:"password" json:"password"  validate:"required,min=6,max=20" example:"123456789"`
	//تکرار پسورد
	AgainPassword string `form:"again_password" json:"again_password"  validate:"required,min=6,max=20,eqfield=Password" example:"123456789"`
}

type ForgetPassword struct {
	//پسورد
	Mobile string `form:"mobile" json:"mobile"  validate:"required,min=11,max=11,startswith=09" example:"09024809750"`
}
