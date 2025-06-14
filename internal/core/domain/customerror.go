package domain

import "errors"

var (
	ErrInvalidBoard    = errors.New("invalid board")
	ErrInvalidSquence  = errors.New("invalid sequence")
	ErrInvalidCol      = errors.New("invalid column")
	ErrColNotAvailable = errors.New("column not available")
)
