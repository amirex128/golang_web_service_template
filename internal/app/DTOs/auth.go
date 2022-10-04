package DTOs

type Verify struct {
	Mobile     string `form:"mobile" json:"mobile" binding:"required" validate:"required,min=11,max=11,startswith=09"`
	VerifyCode string `form:"verify_code" json:"verify_code" validate:"required,min=4,max=4"`
}
type RequestLoginRegister struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required" validate:"required,min=11,max=11,startswith=09"`
}
