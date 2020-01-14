package commons

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fdiezdev/edcomments/models"
)

// DisplayMessage returns a message to the client
func DisplayMessage(w http.ResponseWriter, m models.Message) {

	j, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error al enviar el mensaje: %s", err)
	}

	w.WriteHeader(m.StatusCode)
	w.Write(j)

}
