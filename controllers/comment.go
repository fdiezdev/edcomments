package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/fdiezdev/edcomments/commons"
	"github.com/fdiezdev/edcomments/config"
	"github.com/fdiezdev/edcomments/models"
)

//"github.com/fdiezdev/edcomments/models"

// CreateComment -> creates a comment
func CreateComment(w http.ResponseWriter, r *http.Request) {
	comment := models.Comment{}
	m := models.Message{}
	user := models.User{}

	user, _ = r.Context().Value("user").(models.User)

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		m.StatusCode = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error: %s", err)

		commons.DisplayMessage(w, m)
		return
	}

	comment.UserID = user.ID

	db := config.GetConn()
	defer db.Close()

	err = db.Create(&comment).Error

	if err != nil {
		m.StatusCode = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error: %s", err)

		commons.DisplayMessage(w, m)
		return
	}

	m.StatusCode = http.StatusCreated
	m.Message = "Comentario creado"

	commons.DisplayMessage(w, m)

}

// GetComments -> returns all the comments in a GET request
func GetComments(w http.ResponseWriter, r *http.Request) {

	comments := []models.Comment{}
	m := models.Message{}
	user := models.User{}
	vote := models.Vote{}

	user, _ = r.Context().Value(userCtxKey).(models.User)

	vars := r.URL.Query()

	db := config.GetConn()
	defer db.Close()

	query := db.Where("parent_id = 0")

	if order, ok := vars["order"]; ok {

		if order[0] == "votes" {
			query = query.Order("votes desc, created_at desc")
		}

	} else {

		if idlimit, ok := vars["idlimit"]; ok {
			registerByPage := 30
			offset, err := strconv.Atoi(idlimit[0])

			if err != nil {
				log.Println("Error: ", err)
			}

			query = query.Where("id BETWEEN ? AND ?", offset-registerByPage, offset)
		}

		query = query.Order("id desc")

	}

	query.Find(&comments)

	for i := range comments {

		db.Model(&comments[i]).Related(&comments[i].User)
		comments[i].User[0].Password = ""

		comments[i].Children = getCommentChildren(comments[i].ID)

		// Busco el voto del usuario en sesión

		vote.CommentID = comments[i].ID
		vote.UserID = user.ID

		count := db.Where(&vote).Find(&vote).RowsAffected

		if count > 0 {
			if vote.Value {
				comments[i].HasVoted = 1
			} else {
				comments[i].HasVoted = -1
			}
		}

	}

	j, err := json.Marshal(comments)

	if err != nil {
		m.StatusCode = http.StatusInternalServerError
		m.Message = "Error al formatear los comentarios"

		commons.DisplayMessage(w, m)

		return
	}

	if len(comments) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		m.StatusCode = http.StatusNoContent
		m.Message = "No hay comentarios con las variables aplicadas"

		commons.DisplayMessage(w, m)
	}

}

func getCommentChildren(id uint) (children []models.Comment) {

	db := config.GetConn()
	defer db.Close()

	db.Where("parent_id = ?", id).Find(&children)

	for i := range children {

		db.Model(&children[i]).Related(&children[i].User)
		children[i].User[0].Password = ""

	}

	return

}
