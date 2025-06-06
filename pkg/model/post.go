package model

import (
	"github.com/google/uuid"
	"time"
)

func NewPost(title, short, content string) *Post {
	return &Post{
		Id:        uuid.New(),
		Title:     title,
		ShortDesc: short,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type Post struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	ShortDesc string    `json:"shortDescription"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ViewCount int       `json:"viewCount"`
}

func (p *Post) Filename() string {
	return p.Id.String() + ".json"
}
