package entity

type Contragent struct {
	Id      int64  `json:"id" db:"id"`
	Name    string `json:"name" db:"name" binding:"required"`
	Address string `json:"address" db:"address" binding:"required"`
	InnKpp  string `json:"inn_kpp" db:"innkpp" binding:"required,min=20,max=20"`
}

type UpdateContragentInput struct {
	Name    *string `json:"name"`
	Address *string `json:"address"`
	InnKpp  *string `json:"inn_kpp"`
}

