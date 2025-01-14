package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/quanergyo/ozon-test-assingment/graph/model"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, userID string, title string, content string, commentsEnabled bool) (*model.Post, error) {
	id, err := strconv.Atoi(userID)
	if err != nil {
		return nil, err
	}
	slog.Info("CreatePost resolver")
	data, err := r.repo.CreatePost(id, title, content, commentsEnabled)
	if err != nil {
		slog.Info("Error to write in DB:", err.Error())
	}
	return data, err
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, postID *string, title *string, content *string, commentsEnabled *bool) (*model.Post, error) {
	id, err := strconv.Atoi(*postID)
	if err != nil {
		return nil, err
	}

	return r.repo.UpdatePost(id, title, content, commentsEnabled)
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, postID *string) (*model.Answer, error) {
	id, err := strconv.Atoi(*postID)
	if err != nil {
		return nil, err
	}

	return &model.Answer{Answer: "Ok"}, r.repo.DeletePost(id)
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, userID string, postID string, parentID *string, content string) (*model.Comment, error) {
	if len(content) > 2000 {
		return nil, fmt.Errorf("content must be 2000 charactes or less")
	}
	userId, err := strconv.Atoi(userID)
	if err != nil {
		return nil, err
	}
	postId, err := strconv.Atoi(postID)
	if err != nil {
		return nil, err
	}

	return r.repo.CreateComment(userId, postId, parentID, content)
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	return r.repo.GetAllPosts()
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	postId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return r.repo.GetPost(postId)
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context, postID string, page *int) ([]*model.Comment, error) {
	id, err := strconv.Atoi(postID)
	if err != nil {
		return nil, err
	}
	if page != nil {
		return r.repo.GetCommentsByPage(id, *page)
	}
	return r.repo.GetAllComments(id)
}

// CommentAdded is the resolver for the commentAdded field.
func (r *subscriptionResolver) CommentAdded(ctx context.Context, postID string) (<-chan *model.Comment, error) {
	data, err := r.Query().Comments(ctx, postID, nil)
	if err != nil {
		return nil, err
	}
	commentsChannel := make(chan *model.Comment, len(data))
	for _, item := range data {
		commentsChannel <- item
	}
	close(commentsChannel)
	return commentsChannel, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
