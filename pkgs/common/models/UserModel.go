package models

type UserModel struct {
	ID int64
	Name string
	Age uint
	CreatedAt int64
	UpdatedAt int64
}

func (user *UserModel) TableName() string {
	return "cp_user"
}