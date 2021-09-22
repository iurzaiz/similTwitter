package bd

import (
	"context"
	"time"

	"github.com/iurzaiz/similTwitter/models"
)

//InsertoRelacion graba la relacion en la BD
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("similTwitterDB")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
