package repositories

import (
	"strings"

	"github.com/codingwithik/simple-blog-backend-with-go/src/config"
	"github.com/codingwithik/simple-blog-backend-with-go/src/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *models.User) error
	FindById(id string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	DoesEmailExist(email string) (bool, error)
	DeleteById(id string) error
	FindAll() *[]models.User
	WithTx(tx *gorm.DB) IUserRepository
}

type userRepo struct {
	db *gorm.DB
}

// NewUserRepo will instantiate User Repository
func NewUserRepo() IUserRepository {
	return &userRepo{
		db: config.DB(),
	}
}

func (r *userRepo) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) FindById(id string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User

	if err := r.db.Unscoped().Where("LOWER(email) = ?", strings.ToLower(email)).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) FindAll() *[]models.User {
	var users []models.User
	r.db.Find(&users)
	return &users
}

func (r *userRepo) DoesEmailExist(email string) (bool, error) {
	var user models.User

	if err := r.db.Unscoped().Where("LOWER(email) = ?", strings.ToLower(email)).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *userRepo) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepo) DeleteById(id string) error {
	if err := r.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepo) WithTx(tx *gorm.DB) IUserRepository {
	return &userRepo{db: tx}
}
