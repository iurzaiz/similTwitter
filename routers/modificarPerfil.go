package routers

import (
	"encoding/json"
	"net/http"

	"github.com/iurzaiz/similTwitter/bd"
	"github.com/iurzaiz/similTwitter/models"
)

//ModificarPerfil
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "datos incorrectos "+err.Error(), 400)
		return
	}
	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "ocurrio un error al intenttar modificar el registtro, reintente nuevamente"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado modificar el registro del usuario"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
