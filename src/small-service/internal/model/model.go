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
type CardInfo struct {
	Wid int
	Wsid string
	CardName string
	CardCat string
	Color string
	Prop string
	Rare string
	Level int
	Cost int
	Judge int
	Soul int
	Attack int
	Series string
	Des1 string
	Des2 string
	Des3 string
	Cover1 string
	Cover2 string
	Cover3 string
	Rel1 string
	Rel2 string
	Cat string
}