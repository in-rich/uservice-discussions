package services

import "errors"

var (
	ErrInvalidData              = errors.New("invalid data")
	ErrMessageInWrongDiscussion = errors.New("message in wrong discussion")
)
