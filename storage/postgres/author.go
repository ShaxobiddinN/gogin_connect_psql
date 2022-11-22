package postgres

import (
	"errors"
	"http-server/models"
)

//AddAuthor...
func (stg Postgres) AddAuthor(id string, entity models.CreateAuthorModel) error {
	
	
	_,err := stg.db.Exec(`INSERT INTO author 
	(id,firstname, lastname) 
	VALUES ($1, $2, $3)`, 
	id,
	entity.Firstname,
	entity.Lastname,
)
	if err != nil {
		return err
	}
	
/* 	var author models.Author

	author.Id = id
	author.Firstname=entity.Firstname
	author.Lastname=entity.Lastname
	author.CreatedAt = time.Now()

	im.Db.InMemoryAuthorData = append(im.Db.InMemoryAuthorData, author) */

	return nil
}

//GetAuthorById...
func (stg Postgres) GetAuthorById(id string) (models.Author, error) {
	var a models.Author
	err := stg.db.QueryRow(`SELECT 
	id, 
	firstname, 
	lastname, 
	created_at, 
	updated_at, 
	deleted_at
	FROM author
	WHERE id = $1`,id).Scan(
		&a.Id,
		&a.Firstname,
		&a.Lastname,
		&a.CreatedAt,
		&a.UpdateAt,
		&a.DeleteAt,


	)
	if err != nil {
		return a,err
	}

	 	// var result models.Author
	/*for _, v := range im.Db.InMemoryAuthorData {
		if v.Id == id && v.DeleteAt == nil {
			result = v
			return result, nil
		}
	} */
	return a, nil
}

//GetAuthorList...	
func (stg Postgres) GetAuthorList(offset,limit int,search string) (resp []models.Author,err error){
	rows, err := stg.db.Queryx(`SELECT 
	id, 
	firstname, 
	lastname,  
	created_at, 
	updated_at, 
	deleted_at 
	FROM author 
	WHERE deleted_at IS NULL AND (firstname ILIKE '%' || $1 || '%') OR (lastname ILIKE '%' || $1 || '%')
	LIMIT $2
	OFFSET $3`, search,limit,offset)
	if err != nil {
		return resp,err
	}


	for rows.Next() {
		var a models.Author

		err := rows.Scan(
			&a.Id,
			&a.Firstname,
			&a.Lastname,
			&a.CreatedAt,
			&a.UpdateAt,
			&a.DeleteAt,
		)
		if err != nil {
			return resp,err
		}
		resp = append(resp, a)
	}	
	return resp, err
}

//UpdateAuthor...
func (stg Postgres) UpdateAuthor(entity models.UpdateAuthorModel) error{
	res, err := stg.db.NamedExec(`
	UPDATE  author SET 
		firstname=:f, 
		lastname=:l,
		updated_at=now() 
		WHERE id=:i AND deleted_at IS NULL `, map[string]interface{}{
		"f": entity.Firstname,
		"l": entity.Lastname,
		"i": entity.Id,
	})
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n > 0 {
		return nil
	}
	/* for i, v := range im.Db.InMemoryAuthorData {
		if v.Id == entity.Id && v.DeleteAt==nil{
			v.Firstname = entity.Firstname
			v.Lastname = entity.Lastname

			t := time.Now()
			v.UpdateAt = &t
			im.Db.InMemoryAuthorData[i] = v
			
			return nil
		}
	} */
	return errors.New("author not found")
}

//RemoveAuthor...
func (stg Postgres) RemoveAuthor(id string) error{
	res, err := stg.db.Exec(`UPDATE author SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL`, id)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if  err != nil {
		return err
	}
	if n>0 {
		return nil
	}
	/* for i, v := range im.Db.InMemoryAuthorData {
		if v.Id == id && v.DeleteAt==nil {
			if v.DeleteAt!=nil{
				return errors.New("already deleted")
			}
			t:=time.Now()
			v.DeleteAt=&t
			im.Db.InMemoryAuthorData[i] = v
			return nil
		}
	} */
	return errors.New("author not found or already deleted")
}