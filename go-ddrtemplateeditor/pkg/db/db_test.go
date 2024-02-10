package db

import (
	"bytes"
	"image"
	"image/png"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCRUDOperations(t *testing.T) {
	// Create a new database
	db, err := NewDB()
	assert.NoError(t, err)
	assert.NotNil(t, db)

	// Insert a template
	name := "Test Template"
	img := getImage()
	id, err := db.CreateTemplate(&Template{Name: name, Img: img})
	assert.NoError(t, err)
	assert.NotZero(t, id)

	// Query the inserted template
	template, err := db.QueryTemplate(id)
	assert.NoError(t, err)
	assert.NotNil(t, template)
	assert.Equal(t, name, template.Name)
	assert.Equal(t, img, template.Img)

	// Query all templates
	templates, err := db.QueryTemplates()
	assert.NoError(t, err)
	assert.NotNil(t, templates)
	assert.Len(t, templates, 1)
	assert.Equal(t, name, templates[0].Name)
	assert.Equal(t, img, templates[0].Img)

	// Drop the database
	err = db.DropDb()
	assert.NoError(t, err)

	// Remove the database file
	err = os.Remove("ddrtemplateeditor.db")
	assert.NoError(t, err)
}

func getImage() []byte {
	file, err := os.Open("../../samples/templates/template1.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	var buffer bytes.Buffer
	err = png.Encode(&buffer, img)
	if err != nil {
		log.Fatal(err)
	}

	return buffer.Bytes()
}
