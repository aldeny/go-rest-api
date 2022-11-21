package models

type Book struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	BookName    string `gorm:"type:varchar(225)" json:"book_name"`
	Description string `gorm:"type:text" json:"description"`
}
