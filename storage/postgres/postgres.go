package postgres

import (

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

)

var schema = `
CREATE TABLE IF NOT EXISTS article (
	id CHAR(36) PRIMARY KEY,
	title VARCHAR(255) UNIQUE NOT NULL,
	body TEXT NOT NULL,
	author_id CHAR(36),
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);
CREATE TABLE IF NOT EXISTS author (
	id CHAR(36) PRIMARY KEY,
	firstname VARCHAR(255) NOT NULL,
	lastname VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);
ALTER TABLE article DROP CONSTRAINT IF EXISTS fk_article_author;
ALTER TABLE article ADD CONSTRAINT fk_article_author FOREIGN KEY (author_id) REFERENCES author (id);
`

//InMemory...
 type Postgres struct{
	db *sqlx.DB

 }

//InitDb...
func InitDb(psqlConfig string) (*Postgres, error) {
	var err error
	//"user=admin password=postgres dbname=article_db sslmode=disable"
	tempDB, err := sqlx.Connect("postgres", psqlConfig)
	if err != nil {
		return nil, err
	}

	tempDB.MustExec(schema)

	tx := tempDB.MustBegin()

	tx.MustExec("INSERT INTO author (id, firstname, lastname) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING", "3e1dfc06-dcf6-41fc-b3cc-7c0563fdfab3", "John", "Doe")
	tx.MustExec("INSERT INTO author (id, firstname, lastname) VALUES ($3, $2, $1) ON CONFLICT DO NOTHING", "Najmiddinov", "Shaxobiddin", "24000e82-9c48-4297-a442-ecd1ad55791e")

	tx.MustExec("INSERT INTO article (id, title, body, author_id) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING", "26e2aebc-9771-45ba-8577-ef1a2e7b4170", "Lorem 1", "Body 1", "3e1dfc06-dcf6-41fc-b3cc-7c0563fdfab3")
	tx.MustExec("INSERT INTO article (id, title, body, author_id) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING", "9900756f-e3ed-4dd7-a3a8-4e3cef248ccc", "Lorem 2", "Body 2", "24000e82-9c48-4297-a442-ecd1ad55791e")

	tx.NamedExec("INSERT INTO article (id, title, body, author_id) VALUES (:id, :t, :b, :aid) ON CONFLICT DO NOTHING", map[string]interface{}{
		"id":  "3e451dc4-42e8-4dbc-a70b-edee8f6452ba",
		"t":   "Lorem 3",
		"b":   "Body 3",
		"aid": "3e1dfc06-dcf6-41fc-b3cc-7c0563fdfab3",
	})

	err = tx.Commit()
	if err != nil {
		return nil,err
	}


	return 	&Postgres{
		db: tempDB,
	},nil
}



