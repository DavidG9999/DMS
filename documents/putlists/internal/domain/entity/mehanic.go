package entity

type Mehanic struct {
	Id       int64  `json:"id" db:"id"`
	FullName string `json:"full_name" db:"fullname" binding:"required"`
}

type UpdateMehanicInput struct {
	FullName *string `json:"full_name"`
}

