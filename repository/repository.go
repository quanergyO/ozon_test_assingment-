package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/quanergyo/ozon-test-assingment/graph/model"
	"github.com/quanergyo/ozon-test-assingment/repository/inmemory"
	"github.com/quanergyo/ozon-test-assingment/repository/postgres"
)

type User interface {
	// TODO implement user Auth
}

type Post interface {
	CreatePost(id int, title string, content string, commentsEnabled bool) (*model.Post, error)
	UpdatePost(postID int, title *string, content *string, commentsEnabled *bool) (*model.Post, error)
	DeletePost(id int) error
	GetPost(id int) (*model.Post, error)
	GetAllPosts() ([]*model.Post, error)
}

type Comment interface {
	CreateComment(userId int, postID int, parentID *string, content string) (*model.Comment, error)
	GetAllComments(postId int) ([]*model.Comment, error)
	GetCommentsByPage(postId int, page int) ([]*model.Comment, error)
}

type Repository struct {
	Post
	Comment
}

func NewRepository(db *sqlx.DB) *Repository {
	if db == nil {
		return &Repository{
			Post:    inmemory.NewMemoryPost(),
			Comment: inmemory.NewMemoryComment(),
		}
	}
	return &Repository{
		Post:    postgres.NewPost(db),
		Comment: postgres.NewComment(db),
	}
}
