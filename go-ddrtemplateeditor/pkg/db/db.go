package db

import (
	"bytes"
	"database/sql"
	"image"
	"log"
	"models"

	_ "github.com/mattn/go-sqlite3"
)

func getDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./ddrtemplateeditor.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Create() {
	db, err := getDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS templates (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			image BLOB NOT NULL
		)`)
	if err != nil {
		log.Fatal(err)
	}
}

func Insert(name string, image []byte) (int, error) {
	db, err := getDb()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	r := db.QueryRow(`INSERT INTO templates (name, image) VALUES (?, ?) RETURNING id`, name, image)

	var id int
	err = r.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func DropDb() {
	db, err := getDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`DROP DATABASE templates`)
	if err != nil {
		log.Fatal(err)
	}
}

func QueryTemplates() ([]*models.Template, error) {
	db, err := getDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, image FROM templates")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	result := []*models.Template{}
	for rows.Next() {
		var id int
		var name string
		var i []byte
		err := rows.Scan(&id, &name, &i)
		if err != nil {
			return nil, err
		}

		img, _, err := image.Decode(bytes.NewReader(i))
		if err != nil {
			return nil, err
		}

		result = append(result, &models.Template{ID: id, Name: name, Img: img})
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func QueryTemplate(id int) (*models.Template, error) {
	db, err := getDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT id, name, image FROM templates WHERE id = ?", id)
	template := models.Template{}
	b := []byte{}
	err = row.Scan(&template.ID, &template.Name, &b)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	template.Img = img

	return &template, nil
}
