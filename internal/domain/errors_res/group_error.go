package errorsres

import "errors"

var (
	ErrGroupNotFound    = errors.New("group not found")
	ErrGroupNameExist   = errors.New("group name already exists for this edu level")
	ErrInvalidSalePrice = errors.New("invalid sale price")
)
