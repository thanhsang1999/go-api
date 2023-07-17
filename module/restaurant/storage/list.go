package restaurantstorage

import (
	"context"
	"go-api/common"
	restaurantmodel "go-api/module/restaurant/model"
)

func (s *sqlStore) ListDataWithCondition(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("status in (1)")
	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("owner_id = ?", f.OwnerId)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	offset := (paging.Page - 1) * paging.Size

	if err := db.Offset(offset).Limit(paging.Size).Order("id desc").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
