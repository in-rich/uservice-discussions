package dao

import "errors"

var (
	ErrMessageNotFound              = errors.New("message not found")
	ErrDiscussionReadStatusNotFound = errors.New("discussion read status not found")
)
