package usermodel

type User struct {
	Id int `json:"id" gorm:"column:id;"`
}

type UserCreate struct {
	Id int `json:"id" gorm:"column:id;"`
}

func (u *UserCreate) Validate() error {
	return nil
}
