package usermodel

type User struct {
	Id    int    `json:"id" gorm:"column:id"`
	Email string `json:"email" gorm:"column:email"`
}
