package users

const (
	UserRoleReadOnly         userRole = "READONLY"
	UserRoleInvoiceOnly      userRole = "INVOICEONLY"
	UserRoleStandard         userRole = "STANDARD"
	UserRoleFinancialAdviser userRole = "FINANCIALADVISER"
	UserRoleManagedClient    userRole = "MANAGEDCLIENT"  // Partner Edition only
	UserRoleCashbookClient   userRole = "CASHBOOKCLIENT" // Partner Edition only
)
