package inmemory

import (
	"fmt"
	"github.com/quanergyo/ozon-test-assingment/graph/model"
)

type Post struct {
	posts map[int]model.Post
}

func NewMemoryPost() *Post {
	return &Post{
		make(map[int]model.Post),
	}
}

func (r *Post) CreatePost(id int, title string, content string, commentsEnabled bool) (*model.Post, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *Post) UpdatePost(postID int, title *string, content *string, commentsEnabled *bool) (*model.Post, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *Post) DeletePost(id int) error {
	return fmt.Errorf("not implemented")
}

func (r *Post) GetPost(id int) (*model.Post, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *Post) GetAllPosts() ([]*model.Post, error) {
	return nil, fmt.Errorf("not implemented")
}
