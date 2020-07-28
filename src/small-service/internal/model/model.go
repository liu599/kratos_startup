package model

import "time"

// Kratos hello kratos.
type Kratos struct {
	Hello string
}

type Article struct {
	RegId int64
	RegName string
	Author string
}
type Regression struct {
	ID int64
	RegId string
	RegName string
	Description string
	RelativePath string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}