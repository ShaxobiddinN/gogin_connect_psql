package inmemory

import "http-server/models"

//InMemory...
 type InMemory struct{
	Db *DB

 }

 //mock...
 type DB struct{
	//InMemoryArticleData...
	InMemoryArticleData []models.Article

	//InMemoryAuthorData...
	InMemoryAuthorData []models.Author
 }