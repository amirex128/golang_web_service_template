package models

type Order struct {
	ID                 int64  `json:"id"`
	UserID             int64  `json:"user_id"`
	CustomerID         int64  `json:"customer_id"`
	FactorNo           string `json:"factor_no"`
	TotalWeight        string `json:"total_weight"`
	IP                 string `json:"ip"`
	DeliveryType       string `json:"delivery_type"`
	Status             string `json:"status"` // suspend ready wrong_ready seller_not_to_attend personal_sent service_sent virtual_sent posted unacceptable waited not_distribution pre_distribution distributed confirmed accept return return_final cancel khesarat gheramati amadesazi merge inprocessing ready_schedule logistic_sent pre_return_logistic ready_logistic
	LastUpdateStatusAt string `json:"last_update_status_at"`
	CreatedAt          string `json:"created_at"`
}
