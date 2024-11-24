package entity

type Driver struct {
	Id       int64  `json:"id" db:"id"`
	FullName string `json:"full_name" db:"fullname" binding:"required"`
	License  string `json:"license" db:"license" binding:"required,min=10,max=10"`
	Class    string `json:"class" db:"class" binding:"required"`
}

type UpdateDriverInput struct {
	FullName *string `json:"full_name"`
	License  *string `json:"license"`
	Class    *string `json:"class"`
}

