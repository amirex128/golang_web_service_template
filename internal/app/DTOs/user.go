package DTOs

type UpdateUser struct {
	Gender     string `form:"gender" json:"gender" validate:"required,oneof=man woman"`
	Firstname  string `form:"firstname" json:"firstname" validate:"required"`
	Lastname   string `form:"lastname" json:"lastname" validate:"required"`
	Email      string `form:"email" json:"email" validate:"required,email"`
	Mobile     string `form:"mobile" json:"mobile" validate:"required,startswith=09,numeric"`
	CartNumber string `form:"cart_number" json:"cart_number" validate:"required,numeric,min=16,max=16"`
	Shaba      string `form:"shaba" json:"shaba" validate:"required,numeric,min=24,max=24"`
}
