package entity

type PutlistHeader struct {
	Id            int64  `json:"id" db:"id"`
	UserId        int64  `json:"user_id" db:"userid"`
	Number        int64  `json:"number" db:"number" binding:"required"`
	BankAccountId int64  `json:"bank_account_id" db:"bankaccountid" binding:"required"`
	DateWith      string `json:"date_with" db:"datewith" binding:"required"`
	DateFor       string `json:"date_for" db:"datefor" binding:"required"`
	AutoId        int64  `json:"auto_id" db:"autoid" binding:"required"`
	DriverId      int64  `json:"driver_id" db:"driverid" binding:"required"`
	DispetcherId  int64  `json:"dispetcher_id" db:"dispetcherid" binding:"required"`
	MehanicId     int64  `json:"mehanic_id" db:"mehanicid" binding:"required"`
}

type PutlistBody struct {
	Id            int64  `json:"id" db:"id"`
	PutlistNumber int64  `json:"putlist_number" db:"putlistheadernumber"`
	Number        int64  `json:"number" db:"number" binding:"required"`
	ContragentId  int64  `json:"contragent_id" db:"contragentid" binding:"required"`
	Item          string `json:"item" db:"item" binding:"required"`
	TimeWith      string `json:"time_with" db:"timewith" binding:"required"`
	TimeFor       string `json:"time_for" db:"timefor" binding:"required"`
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
