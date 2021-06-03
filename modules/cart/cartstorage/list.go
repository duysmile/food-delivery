package cartstorage

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/cart/cartmodel"
	"context"
	"fmt"
	"time"

	"github.com/btcsuite/btcutil/base58"
)

const timeLayout = "2006-01-02T15:04:05.999999"

func (s *sqlStore) ListItemInCartByUserId(ctx context.Context, userId int, paging *common.Paging, moreKeys ...string) ([]cartmodel.Cart, error) {
	db := s.db.Table(cartmodel.Cart{}.TableName()).Where("status in (?)", 1)

	db = db.Where("user_id = ?", userId)
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(v)))
		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("created_at < ?", timeCreated.Format("2006-01-02 15:04:05"))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	var data []cartmodel.Cart
	if err := db.
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range data {
		data[i].Food.UpdatedAt = nil

		if i == len(data)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", data[i].CreatedAt.Format(timeLayout))))
			paging.NextCursor = cursorStr
		}
	}

	return data, nil
}
