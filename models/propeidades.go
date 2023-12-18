package models

import "time"

type Model struct {
	Id                int
	Altura            int
	Ancho             int
	TamañoFuente      int
	estilo            string
	grosor            string
	separacion        int
	decoracion        string
	transformacion    string
	alineacion        string
	identacion        int
	FechaCreacion     time.Time
	FechaModificacion time.Time
	actaivo           bool
}
