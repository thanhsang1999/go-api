package common

import "log"

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
	CurrentUser      = "CURRENT_USER"
)

const (
	TopicLikeRestaurant    = "TopicLikeRestaurant"
	TopicDisLikeRestaurant = "TopicDisLikeRestaurant"
)

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error", err)
	}
}
