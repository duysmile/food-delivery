package userstorage

import (
	"200lab/food-delivery/modules/user/usermodel"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *usermodel.UserCreate) error {
	err := s.db.Create(data).Error
	return err
}
