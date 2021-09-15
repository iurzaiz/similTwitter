package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/iurzaiz/similTwitter/bd"
	"github.com/iurzaiz/similTwitter/models"
)

//GraboTweet permitee grabar el tw en la base de datos
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}
	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar eel registro, intente nuevamente"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "Nose ha logrado insertar el tweet", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
