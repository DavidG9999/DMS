package entity

import "errors"

type BankAccount struct {
	Id                int64  `json:"id"`
	BankAccountNumber string `json:"bank_account_number" binding:"required,min=20,max=20"`
	BankName          string `json:"bank_name" binding:"required"`
	BankIdNumber      string `json:"bank_id_number" binding:"required,min=9,max=9"`
}

type UpdateBankAccountInput struct {
	BankAccountNumber *string `json:"bank_account_number"`
	BankName          *string `json:"bank_name"`
	BankIdNumber      *string `json:"bank_id_number"`
}

func (i UpdateBankAccountInput) Validate() error {
	if i.BankAccountNumber == nil && i.BankName == nil && i.BankIdNumber == nil {
		return errors.New("update structure has no values")
	}
	if i.BankAccountNumber != nil && len(*i.BankAccountNumber) != 20 {
		return errors.New("invalid field format: account_number")
	}
	if i.BankIdNumber != nil && len(*i.BankIdNumber) != 9 {
		return errors.New("invalid field format:  bank_id_number")
	}
	if i.BankName != nil {
		if *i.BankName == "" {
			return errors.New("update structure has empty values")
		}
	}
	return nil
}
