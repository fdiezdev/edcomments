package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/fdiezdev/edcomments/commons"
	"github.com/fdiezdev/edcomments/config"
	"github.com/fdiezdev/edcomments/models"
)

// VoteRegister -> controller to register a vote
func VoteRegister(w http.ResponseWriter, r *http.Request) {
	vote := models.Vote{}
	user := models.User{}
	currentVote := models.Vote{}
	m := models.Message{}

	user, _ = r.Context().Value("user").(models.User)
	err := json.NewDecoder(r.Body).Decode(&vote)

	if err != nil {
		m.StatusCode = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error: %s", err)

		commons.DisplayMessage(w, m)
	}

	vote.UserID = user.ID

	db := config.GetConn()
	defer db.Close()

	db.Where("comment_id = ? AND user_id = ?", vote.CommentID, vote.UserID).First(&currentVote)

	if currentVote.ID == 0 {
		db.Create(&vote)
		err = updateCommentVotes(vote.CommentID, vote.Value)

		if err != nil {
			m.StatusCode = http.StatusBadRequest
			m.Message = err.Error()

			commons.DisplayMessage(w, m)
			return
		}

		m.StatusCode = http.StatusCreated
		m.Message = "Voto registrado"

		commons.DisplayMessage(w, m)
		return
	} else if currentVote.Value != vote.Value {
		currentVote.Value = vote.Value
		db.Save(&currentVote)
		err = updateCommentVotes(vote.CommentID, vote.Value)
		if err != nil {
			m.StatusCode = http.StatusBadRequest
			m.Message = err.Error()

			commons.DisplayMessage(w, m)
			return
		}

		m.StatusCode = http.StatusOK
		m.Message = "Voto actualizado"

		commons.DisplayMessage(w, m)
		return
	}
	m.StatusCode = http.StatusBadRequest
	m.Message = "Ya votaste este comentario"

	commons.DisplayMessage(w, m)

}

func updateCommentVotes(commentID uint, vote bool) (err error) {

	comment := models.Comment{}

	db := config.GetConn()
	defer db.Close()

	rows := db.First(&comment, commentID).RowsAffected

	if rows > 0 {
		if vote {
			comment.Votes++
		} else {
			comment.Votes--
		}

		db.Save(&comment)
	} else {
		err = errors.New("No se encontro un comentario en la base de datos para asignar el voto")
	}

	return
}
