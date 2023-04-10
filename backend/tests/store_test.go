package backend_tests

import (
	"context"
	"testing"

	"ThothNotes/backend"
)

func TestNotesHandler_CreateNote(t *testing.T) {
	type fields struct {
		ctx   context.Context
		store *backend.Store
	}
	type args struct {
		title    string
		summary  string
		content  string
		important bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   backend.Note
	}{
		{
			name: "TestNotesHandler_CreateNote",
			fields: fields{
				ctx: context.Background(),
				store: backend.NewStore("test.db"),
			},
			args: args{
				title: "Test title",
				summary: "Test summary",
				content: "Test content",
				important: true,
			},
			want: backend.Note{
				ID: 1,
				Title: "Test title",
				Summary: "Test summary",
				Content: "Test content",
				Important: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := backend.NewNotesHandlers(tt.fields.store)
			handler.Startup(tt.fields.ctx)
			if got := handler.CreateNote(tt.args.title, tt.args.summary, tt.args.content, tt.args.important); got != tt.want {
				t.Errorf("NotesHandler.CreateNote() = %v, want %v", got, tt.want)
			}
		})
	}
}
