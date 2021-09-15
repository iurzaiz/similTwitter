package models

//Tweet captura del body el mensaje quee nos llega
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
