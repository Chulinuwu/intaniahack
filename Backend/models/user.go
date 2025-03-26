package models

type User struct {
    ID        uint   `json:"id" gorm:"primaryKey"`
    Email     string `json:"email" gorm:"unique;not null"`
    Username  string `json:"username" gorm:"unique;not null"`
    Password  string `json:"-"` 
}