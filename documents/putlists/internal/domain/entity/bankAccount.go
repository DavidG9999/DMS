package entity

type BankAccount struct {
	Id                int64  `json:"id" db:"id"`
	BankAccountNumber string `json:"bank_account_number" binding:"required,min=20,max=20" db:"bankaccountnumber"`
	BankName          string `json:"bank_name" binding:"required" db:"bankname"`
	BankIdNumber      string `json:"bank_id_number" binding:"required,min=9,max=9" db:"bankidnumber"`
	OrganizationId    int64  `json:"organization_id" db:"organizationid"`
}

type UpdateBankAccountInput struct {
	BankAccountNumber *string `json:"bank_account_number"`
	BankName          *string `json:"bank_name"`
	BankIdNumber      *string `json:"bank_id_number"`
}
