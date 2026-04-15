package models

// User 用户表
type User struct {
	UserId   int    `json:"userid" gorm:"primaryKey;not null;unique"`
	UserName string `json:"username"`
	Password string `json:"password"`
}
