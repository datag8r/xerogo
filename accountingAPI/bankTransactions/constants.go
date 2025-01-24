package banktransactions

const (
	BankTransactionTypeSpend            bankTransactionType = "SPEND" // Supports Updates
	BankTransactionTypeSpendOverpayment bankTransactionType = "SPEND-OVERPAYMENT"
	BankTransactionTypeSpendPrepayment  bankTransactionType = "SPEND-PREPAYMENT"
	BankTransactionTypeSpendTransfer    bankTransactionType = "SPEND-TRANSFER" // not supported via non-GET methods

	BankTransactionTypeReceive            bankTransactionType = "RECEIVE" // Supports Updates
	BankTransactionTypeReceiveOverpayment bankTransactionType = "RECEIVE-OVERPAYMENT"
	BankTransactionTypeReceivePrepayment  bankTransactionType = "RECEIVE-PREPAYMENT"
	BankTransactionTypeReceiveTransfer    bankTransactionType = "RECEIVE-TRANSFER" // not supported via non-GET methods
)

const (
	BankTransactionStatusAuthorised bankTransactionStatus = "AUTHORISED"
	BankTransactionStatusDeleted    bankTransactionStatus = "DELETED"
)
