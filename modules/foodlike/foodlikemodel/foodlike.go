package foodlikemodel

const EntityName = "FoodLike"

type FoodLike struct {
	UserId int `json:"-" gorm:"column:user_id;"`
	FoodId int `json:"-" gorm:"column:food_id;"`
}

type FoodLikeCreate struct {
	UserId int `json:"-" gorm:"column:user_id;primarykey;"`
	FoodId int `json:"-" gorm:"column:food_id;primaryKey;"`
}

func (FoodLike) TableName() string {
	return "food_likes"
}

func (FoodLikeCreate) TableName() string {
	return "food_likes"
}

func (f FoodLikeCreate) GetFoodId() int {
	return f.FoodId
}
