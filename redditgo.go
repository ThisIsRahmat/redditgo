package redditgo

import "github.com/google/uuid"

type Thread struct {
	ID uuid.UUID `db:"id"`
	Title string `db:"title"`
	Description string `db:"description"`
}


type Post struct {
	ID uuid.UUID `db:"id"`
	ThreadID uuid.UUID `db:"id"`
	Title string `db:"title"`
	Content string `db:"content"`
	Votes int `db:"votes"`

}


type Comment struct {
		ID uuid.UUID `db:"id"`
		POstID  uuid.UUID `db:"id"`
		Content string `db:"content"`
		Votes int `db:"votes"`
}


// Interfaces determine the storage and database retrieval 


type ThreadStore interface{
	Thread(id uuid.UUID) (Thread, error)
	Threads() ([]Thread, error)
	// Why are we using pointers? 
	CreateThread(t *Thread) error
	UpdateThread(t *Thread) error
	DeleteThread(id uuid.UUID) error
}

type PostStore interface{
	Post(id uuid.UUID) (Post, error)
	PostsByThread(threadID uuid.UUID) ([]Post, error)
	// Why are we using pointers? 
	CreatePost(t *Post) error
	UpdatePost(t *Post) error
	DeletePost(id uuid.UUID) error
}


type CommentStore interface {
	Comment(id uuid.UUID) (Comment, error)
	CommentsbyPost(postID uuid.UUID) ([]Comment, error)
	// Why are we using pointers? 
	CreateComment(t *Comment) error
	UpdateComment(t *Comment) error
	DeleteComment(id uuid.UUID) error
}

