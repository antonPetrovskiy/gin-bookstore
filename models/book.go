package models

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Score  uint `json:"score"`
	Author string `json:"author"`
	Nick string `json:"nick"`
}
