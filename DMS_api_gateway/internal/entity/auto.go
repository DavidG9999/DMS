package entity

import "errors"

type Auto struct {
	Id          int64  `json:"id"`
	Brand       string `json:"brand" binding:"required"`
	Model       string `json:"model" binding:"required"`
	StateNumber string `json:"state_number" binding:"required,min=8,max=9"`
}

type UpdateAutoInput struct {
	Brand       *string `json:"brand"`
	Model       *string `json:"model"`
	StateNumber *string `json:"state_number"`
}

func (i UpdateAutoInput) Validate() error {
	if i.Brand == nil && i.Model == nil && i.StateNumber == nil {
		return errors.New("update structure has no values")
	}
	if i.StateNumber != nil {
		if len(*i.StateNumber) > 12 || len(*i.StateNumber) < 11 {
			return errors.New("invalid state_number param")
		}
	}
	if i.Brand != nil {
		if *i.Brand == "" {
			return errors.New("update structure has empty values")
		}
	}
	if i.Model != nil {
		if *i.Model == "" {
			return errors.New("update structure has empty values")
		}
	}
	return nil
}
