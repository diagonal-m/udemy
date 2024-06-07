package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

// データベースから指定されたメールアドレスを持つユーザーを検索し、
// そのユーザー情報を引数として渡されたuser変数に格納する
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	// First(user): 検索結果の最初のレコードを取得し、そのデータをuserに格納する
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
