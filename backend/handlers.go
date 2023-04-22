package backend

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
)

type NotesHandler struct {
	ctx   context.Context
	store *Store
}

func NewNotesHandlers(store *Store) *NotesHandler {
	return &NotesHandler{
		store: store,
	}
}

func (handler *NotesHandler) Startup(ctx context.Context) {
	handler.ctx = ctx
	handler.store.Init()
}

func (handler *NotesHandler) GetNotes() []Note {
	log.Info("GetNotes() called")
	handler.store.Reset()
	err := handler.store.GetAll()
	if err != nil {
		log.Error(err)
	}
	return handler.store.GetElements()
}

func (handler *NotesHandler) GetNoteByID(id int) Note {
	log.Info("GetNoteByID() called")
	handler.store.Reset()
	handler.store.Get(&Note{ID: id})
	return handler.store.GetElement()
}

func (handler *NotesHandler) CreateNote(title, summary, content string, important bool) {
	handler.store.Reset()
	handler.store.Create(&Note{
		Title:     title,
		Summary:   summary,
		Content:   content,
		CreatedDate: 	time.Now(),
		Important: important,
	})
}

func (handler *NotesHandler) CreateNoteWithTags(title, summary, content string, important bool, tags []Tag) {
	handler.store.Reset()
	handler.store.Create(&Note{
		Title:     title,
		Summary:   summary,
		Content:   content,
		CreatedDate: 	time.Now(),
		Important: important,
		Tags:      tags,
	})
}

func (handler *NotesHandler) UpdateNoteByID(id int, title, summary, content string, important bool, tags []Tag) Note {
	log.Info("UpdateNoteByID() called")
	handler.store.Reset()
	handler.store.Update(&Note{
		ID:        id,
		Title:     title,
		Summary:   summary,
		Content:   content,
		UpdatedDate: 	time.Now(),
		Important: important,
		Tags:      tags,
	})
	return handler.store.GetElement()
}

func (handler *NotesHandler) DeleteNoteByID(id int) {
	log.Info("DeleteNoteByID() called")
	handler.store.Reset()
	handler.store.Delete(&Note{ID: id})
}
