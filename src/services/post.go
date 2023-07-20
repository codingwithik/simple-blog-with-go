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
	user_id, err := ps.userRepo.FindById(rq.UserID)

	if err != nil {
		return errors.New("User with id: " + rq.UserID + " not found")
	}

	post := models.Post{

		Title:   rq.Title,
		Content: rq.Content,
		Image:   rq.Image,
		UserID:  user_id.ID,
	}

	return &ps.postRepo.Create(&post)

}

func (ps *postService) UpdatePost(rq *models.PostRequest, user *models.User) error {

	post, err := ps.postRepo.FindById(rq.ID)

	if err != nil {
		return err
	}

	user_id, err := ps.userRepo.FindById(rq.UserID)

	if err != nil {
		return errors.New("User with id: " + rq.UserID + " not found")
	}

	post.Title = rq.Title
	post.Content = rq.Content
	post.Image = rq.Image
	post.UserID = user_id.ID

	return &ps.postRepo.Update(post)
}

func (ps *postService) ListAllPosts() *[]models.Post {
	return ps.postRepo.FindAll()
}

func (ps *postService) ListAllPostsByUser(userId string) (*[]models.Post, error) {
	user_id, err := ps.userRepo.FindById(userId)

	if err != nil {
		return nil, err
	}
	return ps.postRepo.FindAllByUserId(user_id.ID)
}
