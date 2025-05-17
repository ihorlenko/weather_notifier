package repositories

import (
	"github.com/ihorlenko/weather_notifier/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) Create(email string) (*models.User, error) {
	user := models.User{
		Email: email,
	}
	result := r.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) GetOrCreate(email string) (*models.User, error) {
	user, err := r.GetByEmail(email)
	if err == nil {
		return user, nil
	}

	if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return r.Create(email)
}
