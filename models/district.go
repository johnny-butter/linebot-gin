package models

type District struct {
	ID       int
	Name     string
	CountyID int
	County   County
}

func (District) TableName() string {
	return "district"
}
