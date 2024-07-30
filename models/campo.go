package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Campo struct {
	Id          int
	Posicion    int
	Nombre      string
	Descripcion string
	DataString  string
	DataBinary  primitive.Binary
}
