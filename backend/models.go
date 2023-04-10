package backend

import "time"

type Note struct {
	ID          int       `json:"id" gorm:"auto_increment;primary_key"`
	Title       string    `json:"title" gorm:"type:varchar(100);not null;unique_index"`
	Summary     string    `json:"summary" gorm:"type:varchar(500)"`
	Content     string    `json:"content" gorm:"type:text"`
	Tags        []Tag     `json:"tags" gorm:"many2many:note_tags"` // NoteTags is the join table
	CreatedDate time.Time `json:"created_date" gorm:"<-:create;type:datetime;not null"`
	UpdatedDate time.Time `json:"updated_date" gorm:"type:datetime;not null"`
	Important   bool      `json:"important" gorm:"type:bool"`
}

type Tag struct {
	ID   int    `json:"id" gorm:"auto_increment;primary_key"`
	Name string `json:"name" gorm:"type:varchar(100);not null;unique_index"`
}
