package models

type Note struct {
	ID    *int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Judul *string `json:"judul"`
	Notes *string `json:"notes"`
}
