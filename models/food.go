package models

type Food struct {
	ID          int
	Name        string
	NameEng     string
	Category    string
	CommonNames string `gorm:"type:text"`
	Code        string
}
