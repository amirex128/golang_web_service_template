package constants

type IModel interface {
	GetID() uint64
}

const (
	PendingPaymentOrderStatus        = "pending_payment"
	PendingPaymentOrderPaymentStatus = "pending_payment"

	OKProductStatus      = "ok"
	PendingProductStatus = "pending"
	BlockProductStatus   = "block"

	Suspend           = "suspend"
	Ready             = "ready"
	WrongReady        = "wrong_ready"
	SellerNotToAttend = "seller_not_to_attend"
	PersonalSent      = "personal_sent"
	ServiceSent       = "service_sent"
	VirtualSent       = "virtual_sent"
	Posted            = "posted"
	Unacceptable      = "unacceptable"
	Waited            = "waited"
	NotDistribution   = "not_distribution"
	PreDistribution   = "pre_distribution"
	Distributed       = "distributed"
	Confirmed         = "confirmed"
	Accept            = "accept"
	ReturnOrder       = "return_order"
	ReturnFinal       = "return_final"
	Cancel            = "cancel"
	Khesarat          = "khesarat"
	Gheramati         = "gheramati"
	Amadesazi         = "amadesazi"
	Merge             = "merge"
	Inprocessing      = "inprocessing"
	ReadySchedule     = "ready_schedule"
	LogisticSent      = "logistic_sent"
	PreReturnLogistic = "pre_return_logistic"
	ReadyLogistic     = "ready_logistic"
)
