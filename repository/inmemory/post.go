package inmemory

import (
	"fmt"
	"github.com/quanergyo/ozon-test-assingment/graph/model"
	"strconv"
)

const commentsPerPage = 10

type Post struct {
	count int
	posts map[int]model.Post
}

func NewMemoryPost() *Post {
	return &Post{
		count: 0,
		posts: make(map[int]model.Post),
	}
}

func (r *Post) CreatePost(id int, title string, content string, commentsEnabled bool) (*model.Post, error) {
	r.count++
	r.posts[r.count] = model.Post{
		ID:              strconv.Itoa(r.count),
		UserID:          strconv.Itoa(id),
		Title:           title,
		Content:         content,
		CommentsEnabled: commentsEnabled,
		Comments:        nil,
	}
	data := r.posts[r.count]
	return &data, nil
}

func (r *Post) UpdatePost(postID int, title *string, content *string, commentsEnabled *bool) (*model.Post, error) {
	data := r.posts[postID]
	if title != nil {
		data.Title = *title
	}

	if content != nil {
		data.Content = *content
	}

	if commentsEnabled != nil {
		data.CommentsEnabled = *commentsEnabled
	}
	r.posts[postID] = data
	return &data, nil
}

func (r *Post) DeletePost(id int) error {
	delete(r.posts, id)
	return nil
}

func (r *Post) GetPost(id int) (*model.Post, error) {
	if _, exists := r.posts[id]; !exists {
		return nil, fmt.Errorf("no elements with this id")
	}
	data := r.posts[id]
	return &data, nil
}

func (r *Post) GetAllPosts() ([]*model.Post, error) {
	posts := make([]*model.Post, 0, r.count)
	for _, post := range r.posts {
		p := post
		posts = append(posts, &p)
	}
	return posts, nil

}

func (r *Post) CreateComment(userId int, postID int, parentID *string, content string) (*model.Comment, error) {
	if _, exists := r.posts[postID]; !exists {
		return nil, fmt.Errorf("no post with this id")
	}
	post := r.posts[postID]
	if post.CommentsEnabled == false {
		return nil, fmt.Errorf("comments is not enable")
	}

	comment := &model.Comment{
		ID:       strconv.Itoa(len(post.Comments) + 1),
		UserID:   strconv.Itoa(userId),
		PostID:   strconv.Itoa(postID),
		ParentID: parentID,
		Content:  content,
	}
	if post.Comments == nil {
		post.Comments = make([]*model.Comment, 0, 1)
	}
	post.Comments = append(post.Comments, comment)
	r.posts[postID] = post

	return comment, nil
}

func (r *Post) GetAllComments(postId int) ([]*model.Comment, error) {
	if _, exists := r.posts[postId]; !exists {
		return nil, fmt.Errorf("no post with this id")
	}
	comments := make([]*model.Comment, 0, len(r.posts[postId].Comments))
	for _, comment := range r.posts[postId].Comments {
		c := comment
		comments = append(comments, c)
	}

	return comments, nil
}

func (r *Post) GetCommentsByPage(postId int, page int) ([]*model.Comment, error) {
	if _, exists := r.posts[postId]; !exists {
		return nil, fmt.Errorf("no post with this id")
	}
	idx := page*commentsPerPage - commentsPerPage
	if len(r.posts[postId].Comments) >= idx {
		return nil, fmt.Errorf("no page")
	}

	comments := make([]*model.Comment, 0, commentsPerPage)
	for i := idx; i < idx+commentsPerPage; i++ {
		comment := r.posts[postId].Comments[i]
		comments = append(comments, comment)
	}

	return comments, nil
}
