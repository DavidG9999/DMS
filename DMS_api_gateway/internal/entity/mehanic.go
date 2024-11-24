package entity

import "errors"

type Mehanic struct {
	Id       int64  `json:"id"`
	FullName string `json:"full_name" binding:"required"`
}

type UpdateMehanicInput struct {
	FullName *string `json:"full_name"`
}

func (i UpdateMehanicInput) Validate() error {
	if i.FullName == nil {
		return errors.New("update structure has no values")
	}
	if i.FullName != nil {
		if *i.FullName == "" {
			return errors.New("update structure has empty values")
		}
	}
	return nil
}
