package inmemory

import (
	"fmt"
	"github.com/quanergyo/ozon-test-assingment/graph/model"
)

type Comment struct {
	comments map[int]model.Comment
}

func NewMemoryComment() *Comment {
	return &Comment{
		make(map[int]model.Comment),
	}
}

func (r *Comment) CreateComment(userId int, postID int, parentID *string, content string) (*model.Comment, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *Comment) GetAllComments(postId int) ([]*model.Comment, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *Comment) GetCommentsByPage(postId int, page int) ([]*model.Comment, error) {
	return nil, fmt.Errorf("not implemented")
}
