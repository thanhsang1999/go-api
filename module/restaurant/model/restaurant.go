package restaurantmodel

import (
	"errors"
	"go-api/common"
	"go-api/module/user/usermodel"
	"strings"
)

const EntityName = "restaurants"

type Restaurant struct {
	common.SQLModel
	Name  string          `json:"name" gorm:"column:name"`
	Addr  string          `json:"addr" gorm:"column:addr"`
	Logo  *common.Image   `json:"logo" gorm:"column:logo"`
	Cover *common.Images  `json:"cover" gorm:"column:cover"`
	User  *usermodel.User `json:"user" gorm:"column:user"`
}

func (Restaurant) TableName() string {
	return EntityName
}
func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}

type RestaurantUpdate struct {
	Name  string         `json:"name" gorm:"column:name"`
	Addr  string         `json:"addr" gorm:"column:addr"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo"`
	Cover *common.Images `json:"cover" gorm:"column:cover"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	OwnerId         int            `json:"-" gorm:"column:owner_id"`
	Name            string         `json:"name" gorm:"column:name"`
	Addr            string         `json:"addr" gorm:"column:addr"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo"`
	Cover           *common.Images `json:"cover" gorm:"column:cover"`
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

func (r *RestaurantCreate) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}
func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
