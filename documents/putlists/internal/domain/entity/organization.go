package entity

type Organization struct {
	Id             int64  `json:"id" db:"id"`
	Name           string `json:"name" db:"name" binding:"required"`
	Address        string `json:"address" db:"address" binding:"required"`
	Chief          string `json:"chief" db:"chief" binding:"required"`
	FinancialChief string `json:"financial_chief" db:"financialchief" binding:"required"`
	InnKpp         string `json:"inn_kpp" db:"innkpp" binding:"required,min=20,max=20"`
}

type UpdateOrganizationInput struct {
	Name           *string `json:"name"`
	Address        *string `json:"address"`
	Chief          *string `json:"chief"`
	FinancialChief *string `json:"financial_chief"`
	InnKpp         *string `json:"inn_kpp"`
}

