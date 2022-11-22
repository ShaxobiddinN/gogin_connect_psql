package storage

import (
	"errors"
	"http-server/models"
	"strings"

	// "net/http"
	"time"
)

var InMemoryArticleData []models.Article

func AddArticle(id string, entity models.CreateArticleModel) error {
	var article models.Article
	article.ID = id
	article.Content = entity.Content
	//check author
	article.AuthorID = entity.AuthorID
	article.CreatedAt = time.Now()

	InMemoryArticleData = append(InMemoryArticleData, article)

	return nil
}

func GetArticleById(id string) (models.PackedArticleModel, error) {

	var result models.PackedArticleModel
	for _, v := range InMemoryArticleData {
		if v.ID == id && v.DeleteAt != nil{
			return models.PackedArticleModel{}, errors.New("article already deleted")
		}
		if v.ID == id && v.DeleteAt == nil {
			author, err := GetAuthorById(v.AuthorID)
			if err != nil {
				return result, err
			}

			result.ID = v.ID
			result.Content = v.Content
			result.Author = author
			result.CreatedAt = v.CreatedAt
			result.UpdateAt = v.UpdateAt
			result.DeleteAt = v.DeleteAt
			return result, nil
		}
	}
	return models.PackedArticleModel{}, errors.New("article not found")
}

func GetArticleList(offset, limit int, search string) (resp []models.Article, err error) {
	off := 0
	c := 0
	//delete bolgan yoki bolmaganligi filtrlanvaotti
	for _, v := range InMemoryArticleData {
		if v.DeleteAt == nil && (strings.Contains(v.Title, search) || strings.Contains(v.Body, search)) {
			if offset <= off {
				c++
				resp = append(resp, v)
			}
			if c >= limit {
				break
			}
			c++
		}

	}
	return resp, err
}

func UpdateArticle(entity models.UpdateArticleModel) error {
	for i, v := range InMemoryArticleData {
		if v.ID == entity.ID && v.DeleteAt == nil {
			v.Content = entity.Content

			t := time.Now()
			v.UpdateAt = &t

			InMemoryArticleData[i] = v

			return nil
		}
	}
	return errors.New("article not found")
}

func RemoveArticle(id string) error {
	for i, v := range InMemoryArticleData {
		if v.ID == id && v.DeleteAt == nil {
			t := time.Now()
			v.DeleteAt = &t
			InMemoryArticleData[i] = v
			return nil
		}
	}
	return errors.New("article not found or already deleted")

}
