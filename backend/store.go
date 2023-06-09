package backend

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Store struct {
	dbPath string
	db *gorm.DB
	currentElements *[]Note
	currentElement *Note
}

func NewStore(dbPath string) *Store {
	return &Store{
		dbPath: dbPath,
		currentElements: &[]Note{},
		currentElement: &Note{},
	}
}

func (store *Store) GetElements() []Note {
	return *store.currentElements
}

func (store *Store) GetElement() Note {
	return *store.currentElement
}

func (store *Store) Reset() {
}

func (store *Store) Init() (*Store, error) {
	db, err := gorm.Open(sqlite.Open(store.dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Note{})

	store.db = db

	return store, nil
}

func (store *Store) Get(element *Note) error {
	response := store.db.First(element)

	if response.Error != nil {
		return response.Error
	}

	store.currentElement = element

	return nil
}

func (store *Store) GetAll() error {
	response := store.db.Preload("Tags").Find(store.currentElements)

	if response.Error != nil {
		return response.Error
	}

	return nil
}

func (store *Store) Create(element *Note) error {
	element.CreatedDate = element.CreatedDate.UTC()
	response := store.db.Create(element)

	if response.Error != nil {
		return response.Error
	}

	return nil
}

func (store *Store) Update(element *Note) error {
	original := &Note{
		ID: element.ID,
		Title: element.Title,
		Summary: element.Summary,
		Content: element.Content,
		CreatedDate: element.CreatedDate,
		Important: element.Important,
		Tags: element.Tags,
	}
	err := store.Delete(element)

	if err != nil {
		return err
	}

	err = store.Create(original)

	if err != nil {
		return err
	}

	return nil
}

func (store *Store) Delete(element *Note) error {
	err := store.db.Model(element).Association("Tags").Clear()
	if err != nil {
		return err
	}
	response := store.db.Delete(element)


	if response.Error != nil {
		return response.Error
	}

	return nil
}
