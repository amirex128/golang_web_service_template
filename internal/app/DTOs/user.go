package DTOs

type UpdateUser struct {
	Gender        string `form:"gender" json:"gender" validate:"omitempty,oneof=man woman"`
	Firstname     string `form:"firstname" json:"firstname" validate:"omitempty"`
	Lastname      string `form:"lastname" json:"lastname" validate:"omitempty"`
	Email         string `form:"email" json:"email" validate:"omitempty,email"`
	Mobile        string `form:"mobile" json:"mobile" validate:"omitempty,startswith=09,numeric"`
	CartNumber    string `form:"cart_number" json:"cart_number" validate:"omitempty,numeric,min=16,max=16"`
	Shaba         string `form:"shaba" json:"shaba" validate:"omitempty,numeric,min=24,max=24"`
	Password      string `form:"password" json:"password"  validate:"omitempty,min=6,max=20"`
	AgainPassword string `form:"again_password" json:"again_password"  validate:"omitempty,min=6,max=20,eqfield=Password"`
}
