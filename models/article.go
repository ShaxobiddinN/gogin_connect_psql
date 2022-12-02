package models

import "time"



// Content ...
type Content struct {
	Title string	`json:"title" binding:"required"`
	Body  string	`json:"body" binding:"required"`
}

// Article ...
type Article struct {
	ID          string `json:"id"`
	Content          
	AuthorID    string `json:"author_id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt *time.Time `json:"updated_at"`
	DeleteAt *time.Time `json:"-"`
}	

// CreateArticleModel ...
type CreateArticleModel struct {
	Content          // Promoted fields
	AuthorID    string `json:"author_id" binding:"required"`
}

//PackedArticleModel...
type PackedArticleModel struct {
	ID        string `json:"id"`
	Content          
	GetAuthor    GetAuthor `json:"author"`  //Author    Author `json:"author"` bolgan
	CreatedAt time.Time `json:"created_at"`
	UpdateAt *time.Time `json:"updated_at"`
	DeleteAt *time.Time `json:"d_at"`
}	

//UpdateArticleModel...
type UpdateArticleModel struct {
	ID 			string `json:"id" binding:"required"`
	Content          
}

//Deleted_info...
type Deleted_info struct{
	ID          string `json:"id"`
	Content          
	AuthorID    string `json:"author_id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt *time.Time `json:"updated_at"`
	DeleteAt *time.Time `json:"deleted_at"`
}