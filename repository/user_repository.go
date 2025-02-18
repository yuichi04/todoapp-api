package repository

import (
	"fmt"
	"todoapp-api/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	GetUserById(user *model.User, userId uint) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserById(user *model.User, userId uint) error {
	if userId == 0 {
		return fmt.Errorf("invalid user id")
	}

	// Takeの引数について
	// user: 取得したレコードを格納するための変数
	// userId: 検索条件
	if err := ur.db.Take(user, userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("user not found with id: %d", userId)
		}
		return fmt.Errorf("database error: %v", err)
	}
	return nil
}
