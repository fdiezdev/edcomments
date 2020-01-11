package migrations

import (
	"github.com/fdiezdev/edcomments/config"
	"github.com/fdiezdev/edcomments/models"
)

// Migrate -> DB migration
func Migrate() {
	db := config.GetConn()
	defer db.Close()

	db.CreateTable(&models.User{})    // Users table
	db.CreateTable(&models.Comment{}) // Comment table
	db.CreateTable(&models.Vote{})    // Vote table

	db.Model(&models.Vote{}).AddUniqueIndex(
		"comment_id_user_id_unique",
		"comment_id",
		"user_id") // Key unique index

}
