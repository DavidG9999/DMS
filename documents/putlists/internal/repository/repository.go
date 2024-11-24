package repository

import "errors"

var (
	ErrAutoExists            = errors.New("auto already exists")
	ErrAutoNotFound          = errors.New("auto not found")
	ErrBankAccExists         = errors.New("bank account exists")
	ErrBankAccNotFound       = errors.New("bank account not found")
	ErrContragentExists      = errors.New("contragent already exists")
	ErrContragentNotFound    = errors.New("contragent not found")
	ErrDispethcerNotFound    = errors.New("dispetcher not found")
	ErrDriverExists          = errors.New("driver already exists")
	ErrDriverNotFound        = errors.New("driver not found")
	ErrMehanicNotFound       = errors.New("mehanic not found")
	ErrOrganizationExists    = errors.New("organization already exists")
	ErrOrganizationNotFound  = errors.New("organization not found")
	ErrPutlistExists         = errors.New("putlist already exists")
	ErrPutlistNotFound       = errors.New("putlist not found")
	ErrPutlistBodyNotFound   = errors.New("putlist body not found")
	ErrInvalidDateTimeFormat = errors.New("invalid datetime format")
)
