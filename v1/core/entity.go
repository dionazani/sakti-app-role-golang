package core

import "time"

type appRole struct {
	Id          int32
	Code        string
	Description string
	IsActive    int8
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
