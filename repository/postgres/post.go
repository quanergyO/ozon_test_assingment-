package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/quanergyo/ozon-test-assingment/graph/model"
	"log/slog"
	"strconv"
	"strings"
)

type Post struct {
	db *sqlx.DB
}

func NewPost(db *sqlx.DB) *Post {
	return &Post{
		db: db,
	}
}

func (r *Post) CreatePost(id int, title string, content string, commentsEnabled bool) (*model.Post, error) {
	slog.Info("CreatePost Postgres")
	rows := "userId, title, content, commentsEnabled"

	var postId int
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES($1, $2, $3, $4) RETURNING ID", postTable, rows)
	row := r.db.QueryRow(query, id, title, content, commentsEnabled)
	if err := row.Scan(&postId); err != nil {
		return nil, err
	}

	return &model.Post{
		ID:              strconv.Itoa(postId),
		UserID:          strconv.Itoa(id),
		Title:           title,
		Content:         content,
		CommentsEnabled: commentsEnabled,
		Comments:        nil,
	}, nil
}

func (r *Post) UpdatePost(postID int, title *string, content *string, commentsEnabled *bool) (*model.Post, error) {
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if title != nil {
		setValue = append(setValue, fmt.Sprintf("title=$%d", argId))
		args = append(args, *title)
		argId++
	}

	if content != nil {
		setValue = append(setValue, fmt.Sprintf("content=$%d", argId))
		args = append(args, *content)
		argId++
	}

	if commentsEnabled != nil {
		setValue = append(setValue, fmt.Sprintf("commentsEnabled=$%d", argId))
		args = append(args, *commentsEnabled)
		argId++
	}

	setQuery := strings.Join(setValue, ", ")

	args = append(args, postID)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE ID=$%d RETURNING *", postTable, setQuery, argId)
	rows, err := r.db.Query(query, args...)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var post model.Post
	rows.Next()
	err = rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CommentsEnabled)
	if err != nil {
		return nil, err
	}

	return &post, err
}

func (r *Post) DeletePost(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postTable)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *Post) GetPost(id int) (*model.Post, error) {
	var post model.Post
	rows := fmt.Sprintf("id, userid, title, content, commentsenabled")
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id=$1", rows, postTable)
	err := r.db.Get(&post, query, id)
	if err != nil {
		return nil, err
	}

	var comments []*model.Comment
	rows = "id, userid, postid, parentid, content"
	query = fmt.Sprintf("SELECT %s FROM %s WHERE postid=$1", rows, commentsTable)
	if err := r.db.Select(&comments, query, id); err != nil {
		return nil, err
	}
	post.Comments = comments

	return &post, nil

}

func (r *Post) GetAllPosts() ([]*model.Post, error) {
	postDTO := make([]*model.Post, 0)
	rows := "id, userId, title, content, commentsEnabled"
	query := fmt.Sprintf("SELECT %s FROM %s", rows, postTable)
	slog.Info(query)
	if err := r.db.Select(&postDTO, query); err != nil {
		return nil, err
	}

	return postDTO, nil
}
