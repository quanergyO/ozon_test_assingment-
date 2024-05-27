package inmemory

import (
	"testing"
)

func TestCreatePost(t *testing.T) {
	store := NewMemoryPost()

	post, err := store.CreatePost(1, "Test Title", "Test Content", true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if post.Title != "Test Title" {
		t.Errorf("expected title to be 'Test Title', got %s", post.Title)
	}

	if post.Content != "Test Content" {
		t.Errorf("expected content to be 'Test Content', got %s", post.Content)
	}

	if post.CommentsEnabled != true {
		t.Errorf("expected commentsEnabled to be true, got %v", post.CommentsEnabled)
	}

	if post.UserID != "1" {
		t.Errorf("expected userID to be '1', got %s", post.UserID)
	}
}

func TestUpdatePost(t *testing.T) {
	store := NewMemoryPost()

	// Создаем пост, который будем обновлять
	_, err := store.CreatePost(1, "Initial Title", "Initial Content", true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	newTitle := "Updated Title"
	newContent := "Updated Content"
	newCommentsEnabled := false

	updatedPost, err := store.UpdatePost(1, &newTitle, &newContent, &newCommentsEnabled)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updatedPost.Title != newTitle {
		t.Errorf("expected title to be '%s', got %s", newTitle, updatedPost.Title)
	}

	if updatedPost.Content != newContent {
		t.Errorf("expected content to be '%s', got %s", newContent, updatedPost.Content)
	}

	if updatedPost.CommentsEnabled != newCommentsEnabled {
		t.Errorf("expected commentsEnabled to be %v, got %v", newCommentsEnabled, updatedPost.CommentsEnabled)
	}
}

func TestDeletePost(t *testing.T) {
	store := NewMemoryPost()

	_, err := store.CreatePost(1, "Test Title", "Test Content", true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = store.DeletePost(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = store.GetPost(1)
	if err == nil {
		t.Errorf("expected error when getting deleted post, got none")
	}
}

func TestGetPost(t *testing.T) {
	store := NewMemoryPost()

	// Создаем пост
	post, err := store.CreatePost(1, "Test Title", "Test Content", true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Получаем пост по ID
	retrievedPost, err := store.GetPost(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if retrievedPost.ID != post.ID {
		t.Errorf("expected ID to be '%s', got %s", post.ID, retrievedPost.ID)
	}

	if retrievedPost.Title != post.Title {
		t.Errorf("expected title to be '%s', got %s", post.Title, retrievedPost.Title)
	}

	if retrievedPost.Content != post.Content {
		t.Errorf("expected content to be '%s', got %s", post.Content, retrievedPost.Content)
	}

	if retrievedPost.CommentsEnabled != post.CommentsEnabled {
		t.Errorf("expected commentsEnabled to be %v, got %v", post.CommentsEnabled, retrievedPost.CommentsEnabled)
	}
}
