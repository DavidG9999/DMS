package entity

import "errors"

type Putlist struct {
	Id            int64  `json:"id"`
	UserId        int64  `json:"user_id"`
	Number        int64  `json:"number" binding:"required"`
	BankAccountId int64  `json:"bank_account_id" binding:"required"`
	DateWith      string `json:"date_with" binding:"required"`
	DateFor       string `json:"date_for" binding:"required"`
	AutoId        int64  `json:"auto_id" binding:"required"`
	DriverId      int64  `json:"driver_id" binding:"required"`
	DispetcherId  int64  `json:"dispetcher_id" binding:"required"`
	MehanicId     int64  `json:"mehanic_id" binding:"required"`
}

type PutlistBody struct {
	Id            int64  `json:"id"`
	PutlistNumber int64  `json:"putlist_number"`
	Number        int64  `json:"number" binding:"required"`
	ContragentId  int64  `json:"contragent_id" binding:"required"`
	Item          string `json:"item" binding:"required"`
	TimeWith      string `json:"time_with" binding:"required"`
	TimeFor       string `json:"time_for" binding:"required"`
}

type UpdatePutlistHeaderInput struct {
	BankAccountId *int64  `json:"bank_account_id"`
	DateWith      *string `json:"date_with"`
	DateFor       *string `json:"date_for"`
	AutoId        *int64  `json:"auto_id"`
	DriverId      *int64  `json:"driver_id"`
	DispetcherId  *int64  `json:"dispetcher_id"`
	MehanicId     *int64  `json:"mehanic_id"`
}

type UpdatePutlistBodyInput struct {
	Number       *int64  `json:"number"`
	ContragentId *int64  `json:"contragent_id"`
	Item         *string `json:"item"`
	TimeWith     *string `json:"time_with"`
	TimeFor      *string `json:"time_for"`
}

func (i UpdatePutlistHeaderInput) Validate() error {
	if i.BankAccountId == nil && i.DateWith == nil && i.DateFor == nil && i.AutoId == nil && i.DriverId == nil && i.DispetcherId == nil && i.MehanicId == nil {
		return errors.New("update structure has no values")
	}
	if i.BankAccountId != nil {
		if *i.BankAccountId == 0 {
			return errors.New("update structure has empty values")
		}
	}
	if i.DateWith != nil {
		if *i.DateWith == "" {
			return errors.New("update structure has empty values")
		}
	}
	if i.DateFor != nil {
		if *i.DateFor == "" {
			return errors.New("update structure has empty values")
		}
	}
	if i.AutoId != nil {
		if *i.AutoId == 0 {
			return errors.New("update structure has empty values")
		}
	}
	if i.DriverId != nil {
		if *i.DriverId == 0 {
			return errors.New("update structure has empty values")
		}
	}
	if i.DispetcherId != nil {
		if *i.DispetcherId == 0 {
			return errors.New("update structure has empty values")
		}
	}
	if i.MehanicId != nil {
		if *i.MehanicId == 0 {
			return errors.New("update structure has empty values")
		}
	}
	return nil
}

func (i UpdatePutlistBodyInput) Validate() error {
	if i.Number == nil && i.ContragentId == nil && i.Item == nil && i.TimeWith == nil && i.TimeFor == nil {
		return errors.New("update structure has no values")
	}
	if i.Number != nil {
		if *i.Number == 0 {
			return errors.New("update structure has empty values")
		}
	}
	if i.ContragentId != nil {
		if *i.ContragentId == 0 {
			return errors.New("update structure has empty values")
		}
	}
	if i.Item != nil {
		if *i.Item == "" {
			return errors.New("update structure has empty values")
		}
	}
	if i.TimeWith != nil {
		if *i.TimeWith == "" {
			return errors.New("update structure has empty values")
		}
	}
	if i.TimeFor != nil {
		if *i.TimeFor == "" {
			return errors.New("update structure has empty values")
		}
	}
	return nil
}
