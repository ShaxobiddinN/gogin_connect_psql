package postgres

import (
	"errors"
	"http-server/models"
)

// AddAuthor...
func (stg Postgres) AddAuthor(id string, entity models.CreateAuthorModel) error {

	// fname:=entity.Firstname+entity.Lastname
	_, err := stg.db.Exec(`INSERT INTO author 
	(id,fullname) 
	VALUES ($1, $2)`,
		id,
		entity.Fullname,
		// entity.Firstname,
		// entity.Lastname,
		// entity.Middlename,

	)
	if err != nil {
		return err
	}

	return nil
}

// GetAuthorById...
func (stg Postgres) GetAuthorById(id string) (models.Author, error) {
	var a models.Author
	// var tempMiddlename *string

	//firstname,
	//lastname,
	//middlename,
	err := stg.db.QueryRow(`SELECT 
	id, 
	fullname,
	created_at, 
	updated_at, 
	deleted_at
	FROM author
	WHERE id = $1`, id).Scan(
		&a.Id,
		&a.Fullname,
		// &a.Firstname,
		// &a.Lastname,
		// &tempMiddlename,
		&a.CreatedAt,
		&a.UpdateAt,
		&a.DeleteAt,
	)
	if err != nil {
		return a, err
	}
	// if tempMiddlename != nil {
	// 	a.Middlename = *tempMiddlename
	// }

	return a, nil
}

// GetAuthorList...
func (stg Postgres) GetAuthorList(offset, limit int, search string) (resp []models.Author, err error) {
	rows, err := stg.db.Queryx(`SELECT 
	id, 
	fullname,
	created_at, 
	updated_at, 
	deleted_at 
	FROM author 
	WHERE deleted_at IS NULL AND 
	(fullname ILIKE '%' || $1 || '%') 
	LIMIT $2
	OFFSET $3`, search, limit, offset)
	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var a models.Author

		err := rows.Scan(
			&a.Id,
			// &a.Firstname,
			// &a.Lastname,
			// &a.Middlename,
			&a.Fullname,
			&a.CreatedAt,
			&a.UpdateAt,
			&a.DeleteAt,
		)
		if err != nil {
			return resp, err
		}
		resp = append(resp, a)
	}
	return resp, err
}

// UpdateAuthor...
// firstname=:f,
//
//	lastname=:l,
//	middlename =: m,
func (stg Postgres) UpdateAuthor(entity models.UpdateAuthorModel) error {
	res, err := stg.db.NamedExec(`
	UPDATE  author SET 
		fullname =:fn,
		updated_at=now() 
		WHERE id =:i AND deleted_at IS NULL `, map[string]interface{}{
		// "f": entity.Firstname,
		// "l": entity.Lastname,
		// "m": entity.Middlename,
		"fn": entity.Fullname,
		"i":  entity.Id,
	})
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n > 0 {
		return nil
	}

	return errors.New("author not found")
}

// RemoveAuthor...
func (stg Postgres) RemoveAuthor(id string) error {
	res, err := stg.db.Exec(`UPDATE author SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL`, id)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n > 0 {
		return nil
	}

	return errors.New("author not found or already deleted")
}
