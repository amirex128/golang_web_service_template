package DTOs

type Verify struct {
	Mobile     string `form:"mobile" json:"mobile" validate:"required,min=11,max=11,startswith=09"`
	VerifyCode string `form:"verify_code" json:"verify_code" validate:"omitempty,min=4,max=4"`
	Password   string `form:"password" json:"password"  validate:"omitempty,min=6,max=20"`
}
type RequestLoginRegister struct {
	Mobile string `form:"mobile" json:"mobile"  validate:"required,min=11,max=11,startswith=09"`
}

type ChangePassword struct {
	Password      string `form:"password" json:"password"  validate:"required,min=6,max=20"`
	AgainPassword string `form:"again_password" json:"again_password"  validate:"required,min=6,max=20,eqfield=Password"`
}
