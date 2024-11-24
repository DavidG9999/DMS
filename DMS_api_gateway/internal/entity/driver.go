package entity

import "errors"

type Driver struct {
	Id       int64    `json:"id"`
	FullName string `json:"full_name" binding:"required"`
	License  string `json:"license" binding:"required,min=10,max=10"`
	Class    string `json:"class" binding:"required"`
}

type UpdateDriverInput struct {
	FullName *string `json:"full_name"`
	License  *string `json:"license"`
	Class    *string `json:"class"`
}

func (i UpdateDriverInput) Validate() error {
	if i.FullName == nil && i.License == nil && i.Class == nil {
		return errors.New("update structure has no values")
	}
	if i.License != nil {
		if len(*i.License) != 10 {
			return errors.New("invalid license param")
		}
	}
	if i.FullName != nil {
		if *i.FullName == "" {
			return errors.New("update structure has empty values")
		}
	}
	if i.Class != nil {
		if *i.Class == "" {
			return errors.New("update structure has empty values")
		}
	}
	return nil
}
