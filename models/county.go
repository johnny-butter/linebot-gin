package models

type County struct {
	ID    int
	Name  string
	CwbId string
}

func (County) TableName() string {
	return "county"
}
