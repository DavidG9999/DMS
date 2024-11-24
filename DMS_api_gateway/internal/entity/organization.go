package entity

import "errors"

type Organization struct {
	Id             int64    `json:"id"`
	Name           string `json:"name" binding:"required"`
	Address        string `json:"address" binding:"required"`
	Chief          string `json:"chief" binding:"required"`
	FinancialChief string `json:"financial_chief" binding:"required"`
	InnKpp         string `json:"inn_kpp" binding:"required,min=20,max=20"`
}

type UpdateOrganizationInput struct {
	Name           *string `json:"name"`
	Address        *string `json:"address"`
	Chief          *string `json:"chief"`
	FinancialChief *string `json:"financial_chief"`
	InnKpp         *string `json:"inn_kpp"`
}

func (i UpdateOrganizationInput) Validate() error {
	if i.Name == nil && i.Address == nil && i.Chief == nil && i.FinancialChief == nil && i.InnKpp == nil {
		return errors.New("update structure has no values")
	}
	if i.InnKpp != nil {
		if len(*i.InnKpp) != 20 {
			return errors.New("invalid field format: inn_kpp")
		}
	}
	if i.Name != nil {
		if *i.Name == "" {
			return errors.New("update structure has empty values")
		}
	}
	if i.Address != nil {
		if *i.Address == "" {
			return errors.New("update structure has empty values")
		}
	}
	if i.Chief != nil {
		if *i.Chief == "" {
			return errors.New("update structure has empty values")
		}
	}
	if i.FinancialChief != nil {
		if *i.FinancialChief == "" {
			return errors.New("update structure has empty values")
		}
	}
	return nil
}
