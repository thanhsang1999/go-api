package restaurantlikemodel

import "time"

const (
	TableName = "restaurant_likes"
)

type RestaurantLikes struct {
	RestaurantId int        `json:"-" gorm:"restaurant_id"`
	UserId       int        `json:"-" gorm:"user_id"`
	CreatedAt    *time.Time `json:"-" gorm:"created_at"`
}

func (r RestaurantLikes) TableName() string {
	return TableName
}
func (r RestaurantLikes) GetRestaurantId() int {
	return r.RestaurantId
}
