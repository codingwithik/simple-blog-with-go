package models

type Comment struct {
	Base
	Comment     string `gorm:"type:Text;not null" json:"comment,omitempty"`
	CommentedBy string `gorm:"not null" json:"commented_by,omitempty"`
	PostID      string
	Post        Post
}

type CreateCommentRequest struct {
	Comment     string `json:"comment"  binding:"required"`
	CommentedBy string `json:"commented_by" binding:"required"`
	PostID      string `json:"post_id" binding:"required"`
}

type UpdateComment struct {
	Comment     string `json:"comment,omitempty"`
	CommentedBy string `json:"commented_by,omitempty"`
	PostID      string `json:"post_id,omitempty"`
}
