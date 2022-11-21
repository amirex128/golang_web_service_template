package constants

type IModel interface {
	GetID() uint64
}

const (
	PendingPaymentOrderStatus = "pending_payment"
	FailedPaymentOrderStatus  = "failed_payment"

	CanceledOrderStatus = "canceled"

	PendingAcceptOrderStatus      = "pending_accept"
	AcceptedOrderStatus           = "accepted"
	ChooseCourierOrderStatus      = "choose_courier"
	PendingReceivePostOrderStatus = "pending_receive_post"
	ReceivedPostOrderStatus       = "received_post"
	ReceivedCustomerOrderStatus   = "received_customer"

	PendingReturnOrderStatus  = "pending_return"
	AcceptedReturnOrderStatus = "accepted_return"
	RejectedReturnOrderStatus = "rejected_return"

	PendingReceivePostReturnOrderStatus = "pending_receive_return"
	ReceivedPostReturnOrderStatus       = "received_post_return"
	ReceivedOwnerOrderStatus            = "received_owner"

	FinishedOrderStatus = "finished"

	OKProductStatus      = "ok"
	PendingProductStatus = "pending"
	BlockProductStatus   = "block"

	username = ""
	password = ""
	from     = ""
)
