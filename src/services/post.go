package services

import (
	"errors"

	"github.com/codingwithik/simple-blog-backend-with-go/src/models"
	"github.com/codingwithik/simple-blog-backend-with-go/src/repositories"
)

type IPostService interface {
	CreatePost(rq *models.PostRequest, user *models.User) error
	UpdatePost(rq *models.PostRequest, user *models.User) error
	//DeletePosts(ids []string) error
	ListAllPosts() (*[]models.Post, error)
	ListAllPostsByUser(userId string) (*[]models.Post, error)
}

type postService struct {
	postRepo repositories.IPostRepository
	userRepo repositories.IUserRepository
}

// NewPostService will instantiate AuthService
func NewPostService() IPostService {
	return &postService{
		postRepo: repositories.NewPostRepo(),
		userRepo: repositories.NewUserRepo(),
	}
}

func (ps *postService) CreatePost(rq *models.PostRequest, user *models.User) error {

	// validate userId passed for assigned_to and supervisd_by
	userId, err := ps.userRepo.FindById(rq.UserID)

	if err != nil {
		return errors.New("User with id: " + rq.UserID + " not found")
	}

	post := models.Post{

		Title:   rq.Title,
		Content: rq.Content,
		Image:   rq.Image,
		UserID:  userId.ID,
	}

	return ps.postRepo.Create(&post)

}

func (ps *postService) UpdatePost(rq *models.PostRequest, user *models.User) error {

	post, err := ps.postRepo.FindById(rq.ID)

	if err != nil {
		return err
	}

	userId, err := ps.userRepo.FindById(rq.UserID)

	if err != nil {
		return errors.New("User with id: " + rq.UserID + " not found")
	}

	post.Title = rq.Title
	post.Content = rq.Content
	post.Image = rq.Image
	post.UserID = userId.ID

	return ps.postRepo.Update(post)
}

func (ps *postService) ListAllPosts() (*[]models.Post, error) {
	return ps.postRepo.FindAll(), nil
}

func (ps *postService) ListAllPostsByUser(id string) (*[]models.Post, error) {
	userId, err := ps.userRepo.FindById(id)

	if err != nil {
		return nil, err
	}
	return ps.postRepo.FindAllByUserId(userId.ID)
}
