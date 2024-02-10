package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Template struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Img  []byte `json:"-"`
}

type Item struct {
	ID            int    `json:"id"`
	Type          string `json:"type"`
	X, Y          int    `json:"-"`
	Width, Height int    `json:"-"`
}

type DB struct {
	db *gorm.DB
}

func NewDB() (*DB, error) {
	db, err := gorm.Open(sqlite.Open("ddrtemplateeditor.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Template{}, &Item{})
	if err != nil {
		return nil, err
	}

	return &DB{db: db}, nil
}

func (d *DB) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Close()
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) CreateTemplate(template *Template) (int, error) {
	err := d.db.Create(template).Error
	if err != nil {
		return 0, err
	}

	return template.ID, nil
}

func (d *DB) CreateItem(item *Item) error {
	err := d.db.Create(item).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) DropDb() error {
	err := d.db.Migrator().DropTable(&Template{}, &Item{})
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) QueryTemplates() ([]*Template, error) {
	var templates []*Template
	err := d.db.Find(&templates).Error
	if err != nil {
		return nil, err
	}

	return templates, nil
}

func (d *DB) QueryTemplate(id int) (*Template, error) {
	var template Template
	err := d.db.First(&template, id).Error
	if err != nil {
		return nil, err
	}

	return &template, nil
}

func (d *DB) UpdateTemplateImage(id int, image []byte) error {
	err := d.db.Model(&Template{}).Where("id = ?", id).Update("Img", image).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) QueryItems() ([]*Item, error) {
	var items []*Item
	err := d.db.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (d *DB) QueryItem(id int) (*Item, error) {
	var item Item
	err := d.db.First(&item, id).Error
	if err != nil {
		return nil, err
	}

	return &item, nil
}
