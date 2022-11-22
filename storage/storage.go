package storage

import "http-server/models"

type StorageI interface{
	AddArticle(id string, entity models.CreateArticleModel) error
	GetArticleById(id string) (models.PackedArticleModel, error)
	GetArticleList(offset, limit int, search string) (resp []models.Article, err error)
	UpdateArticle(entity models.UpdateArticleModel) error
	RemoveArticle(id string) error

	AddAuthor(id string, entity models.CreateAuthorModel) error
	GetAuthorById(id string) (models.Author, error)
	GetAuthorList(offset,limit int,search string) (resp []models.Author,err error)
	UpdateAuthor(entity models.UpdateAuthorModel) error
	RemoveAuthor(id string) error
}