package repositories

import (
	"github.com/codingwithik/simple-blog-backend-with-go/src/config"
	"github.com/codingwithik/simple-blog-backend-with-go/src/models"
	"gorm.io/gorm"
)

type IPostRepository interface {
	Create(post *models.Post) error
	Update(post *models.Post) error
	DeleteById(id string) error
	DeleteAllById(ids []string) error
	FindById(id string) (*models.Post, error)
	FindAll() *[]models.Post
	FindAllByUserId(userId string) (*[]models.Post, error)
}

type postRepo struct {
	db *gorm.DB
}

// NewPostRepo will instantiate User Repository
func NewPostRepo() IPostRepository {
	return &postRepo{
		db: config.DB(),
	}
}

func (r *postRepo) Create(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepo) Update(post *models.Post) error {
	return r.db.Save(post).Error
}

func (r *postRepo) DeleteById(id string) error {
	if err := r.db.Delete(&models.Post{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *postRepo) DeleteAllById(ids []string) error {
	if err := r.db.Delete(&models.Post{}, ids).Error; err != nil {
		return err
	}
	return nil
}

func (r *postRepo) FindById(id string) (*models.Post, error) {
	var post models.Post
	if err := r.db.Where("id = ?", id).First(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *postRepo) FindAll() *[]models.Post {
	var posts []models.Post
	r.db.Find(&posts)
	return &posts
}

func (r *postRepo) FindAllByUserId(userId string) (*[]models.Post, error) {
	var post []models.Post
	if err := r.db.Where("userid <> ?", userId).Find(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}
