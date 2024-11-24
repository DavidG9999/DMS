package entity

type User struct{
	ID int64
	Name string
	Email string
	PasswordHash []byte
}

type UpdateNameInput struct{
	Name *string
}

type UpdatePasswordInput struct{
	Password *string
}