package models

type Post struct {
	Base
	Title   string `gorm:"uniqueIndex;not null" json:"title,omitempty"`
	Content string `gorm:"type:Text;not null" json:"content,omitempty"`
	Image   string `json:"image,omitempty"`
	UserID  string
}

type PostRequest struct {
	ID      string `json:"-"`
	Title   string `json:"title"  binding:"required"`
	Content string `json:"content" binding:"required"`
	Image   string `json:"image" binding:"required"`
	UserID  string `json:"user_id,omitempty"`
}
