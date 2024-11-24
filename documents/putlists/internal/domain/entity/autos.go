package entity

type Auto struct {
	Id          int64  `json:"id" db:"id"`
	Brand       string `json:"brand" binding:"required" db:"brand"`
	Model       string `json:"model" binding:"required" db:"model"`
	StateNumber string `json:"state_number" binding:"required,min=8,max=9" db:"statenumber"`
}

type UpdateAutoInput struct {
	Brand       *string `json:"brand"`
	Model       *string `json:"model"`
	StateNumber *string `json:"state_number"`
}


