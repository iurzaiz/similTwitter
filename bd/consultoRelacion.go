package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/iurzaiz/similTwitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ConsultoRelacion consulta relaciones entre usuarios en la base de datos
func ConsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("similTwitterDB")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}
	var resultado models.Relacion
	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
