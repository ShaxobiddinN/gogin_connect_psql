package inmemory

import (
	"errors"
	"http-server/models"
	"strings"
	"time"
)


func (im InMemory) AddAuthor(id string, entity models.CreateAuthorModel) error {
	var author models.Author

	author.Id = id
	author.Firstname=entity.Firstname
	author.Lastname=entity.Lastname
	author.CreatedAt = time.Now()

	im.Db.InMemoryAuthorData = append(im.Db.InMemoryAuthorData, author)

	return nil
}

func (im InMemory) GetAuthorById(id string) (models.Author, error) {

	var result models.Author
	for _, v := range im.Db.InMemoryAuthorData {
		if v.Id == id && v.DeleteAt == nil {
			result = v
			return result, nil
		}
	}
	return result, errors.New("author not found")
}

//GetAuthorList...	
func (im InMemory) GetAuthorList(offset,limit int,search string) (resp []models.Author,err error){
	off := 0
	c := 0
	for _, v := range im.Db.InMemoryAuthorData {
		if v.DeleteAt == nil && (strings.Contains(v.Firstname, search) || strings.Contains(v.Id, search)) {
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

//UpdateAuthor...
func (im InMemory) UpdateAuthor(entity models.UpdateAuthorModel) error{
	for i, v := range im.Db.InMemoryAuthorData {
		if v.Id == entity.Id && v.DeleteAt==nil{
			v.Firstname = entity.Firstname
			v.Lastname = entity.Lastname

			t := time.Now()
			v.UpdateAt = &t
			im.Db.InMemoryAuthorData[i] = v
			
			return nil
		}
	}
	return errors.New("author not found")
}

//RemoveAuthor...
func (im InMemory) RemoveAuthor(id string) error{
	for i, v := range im.Db.InMemoryAuthorData {
		if v.Id == id && v.DeleteAt==nil {
			if v.DeleteAt!=nil{
				return errors.New("already deleted")
			}
			t:=time.Now()
			v.DeleteAt=&t
			im.Db.InMemoryAuthorData[i] = v
			return nil
		}
	}
	return errors.New("author not found or already deleted")
}