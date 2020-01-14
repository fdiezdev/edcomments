package controllers

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fdiezdev/edcomments/commons"
	"github.com/fdiezdev/edcomments/config"
	"github.com/fdiezdev/edcomments/models"
)

// Login -> login controller for user auth
func Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}

	db := config.GetConn()
	defer db.Close()

	crypto := sha256.Sum256([]byte(user.Password))
	pwd := fmt.Sprintf("%x", crypto)

	db.Where("email = ? and password = ?", user.Email, pwd).First(&user)
	log.Println(user.ID, pwd)
	if user.ID > 0 {
		user.Password = ""
		token := commons.GenerateJWT(user)
		j, err := json.Marshal(models.Token{Token: token})
		if err != nil {
			log.Fatalf("Error generando JWT: %s", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		m := models.Message{
			Message:    "Error: Correo o contraseña incorrectos",
			StatusCode: http.StatusUnauthorized,
		}

		commons.DisplayMessage(w, m)
	}
}

// SignUp creates a new user
func SignUp(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		m.Message = fmt.Sprintf("Error: %s", err)
		m.StatusCode = http.StatusBadRequest

		commons.DisplayMessage(w, m)
		return
	}

	if user.Password != user.ConfirmPassword {
		m.Message = fmt.Sprint("Error: las contraseñas no coinciden")
		m.StatusCode = http.StatusBadRequest

		commons.DisplayMessage(w, m)
		return
	}

	if len(user.Password) < 8 {
		m.Message = fmt.Sprint("Error: la contraseña no tiene 8 caracteres como mínimo")
		m.StatusCode = http.StatusBadRequest

		commons.DisplayMessage(w, m)
		return
	}

	crypto := sha256.Sum256([]byte(user.Password))
	pwd := fmt.Sprintf("%x", crypto)

	user.Password = pwd

	picmd5 := md5.Sum([]byte(user.Email))
	picstr := fmt.Sprintf("%x", picmd5)

	user.Picture = "https://gravatar.com/avatar/" + picstr + "?s=100"

	// Saving new user in de DB
	db := config.GetConn()
	defer db.Close()

	err = db.Create(&user).Error
	if err != nil {
		m.Message = "Error: ocurrió algo inesperado al guardar el registro"
		m.StatusCode = http.StatusBadRequest
		commons.DisplayMessage(w, m)

		return
	}

	m.Message = "Registro guardado con éxito"
	m.StatusCode = http.StatusCreated

	commons.DisplayMessage(w, m)
}
