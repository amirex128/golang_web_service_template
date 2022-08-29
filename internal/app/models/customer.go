package models

type Customer struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Family     string `json:"family"`
	Mobile     string `json:"mobile"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	PostalCode string `json:"postal_code"`
	Lat        string `json:"lat"`
	Lng        string `json:"lng"`
}
