package models

import "time"

type Post struct {
	Id        int       `json:"id" form:"id"`
	UserId    int       `json:"userId" form:"userId"`
	Title     string    `json:"title" form:"title"`
	Content   string    `json:"content" form:"content"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt"`
}
