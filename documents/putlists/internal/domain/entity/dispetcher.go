package entity

type Dispetcher struct {
	Id       int64  `json:"id" db:"id"`
	FullName string `json:"full_name" binding:"required" db:"fullname"`
}

type UpdateDispetcherInput struct {
	FullName *string `json:"full_name"`
}