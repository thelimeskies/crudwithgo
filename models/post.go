// model/post.go 

package models

import "time"

type Post struct {
	ID uint `json:"id" grom:"primaryKey"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}