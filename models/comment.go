package models

import (
	"github.com/jinzhu/gorm"
)

// Comment -> comment structure
type Comment struct {
	gorm.Model
	UserID   uint      `json:"userId"`
	ParentID uint      `json:"parentId"`
	Votes    int32     `json:"votes"`
	Content  string    `json:"content"`
	HasVoted int8      `json:"hasVoted" gorm:"-"`
	User     []User    `json:"user,omitempty"`
	Children []Comment `json:"children,omitempty"`
}
