package model

import "time"

func (f *User) TableName() string {
	return "user"
}

type User struct {
	Id        int       `gorm:"primaryKey;autoIncrement:true;column:id;" json:"id"`
	Password  string    `gorm:"column:password" json:"password"`
	Username  string    `gorm:"column:username" json:"username"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}
