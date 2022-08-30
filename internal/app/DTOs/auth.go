package DTOs

type Login struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required" validate:"required,min=11,max=11,startswith=09"`
	Password string `form:"password" json:"password" binding:"required" validate:"required,min=8,max=20"`
}
