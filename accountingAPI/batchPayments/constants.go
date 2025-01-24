package batchpayments

const (
	BatchPaymentTypeBills batchPaymentType = "PAYBATCH"
	BatchPaymentTypeSales batchPaymentType = "RECBATCH"
)

const (
	BatchPaymentStatusAuthorised batchPaymentStatus = "AUTHORISED"
	BatchPaymentStatusDeleted    batchPaymentStatus = "DELETED"
)
