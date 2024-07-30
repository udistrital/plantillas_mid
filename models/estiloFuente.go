package models

type EstiloFuente struct {
	Id                int
	Tamaño            int
	Estilo            string
	Grosor            string
	Altura            int
	Separacion        int
	Decoracion        string
	Transformacion    string
	Alineacion        string
	Identacion        int
	FechaCreacion     string
	FechaModificacion string
	Activo            bool
}
