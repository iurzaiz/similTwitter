package routers

import (
	"encoding/json"
	"net/http"

	"github.com/iurzaiz/similTwitter/bd"
	"github.com/iurzaiz/similTwitter/models"
)

//ConsultaRelacion chequea si hay relacion entre dos usuarios
func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var resp models.RespuestaconsultaRelacion
	status, err := bd.ConsultoRelacion(t)
	if err != nil || !status {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}
