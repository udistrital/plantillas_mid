package models

type Seccion struct {
	//Nuevo modelo seccion
	id                 string
	padre_id           string
	tipo_contenido_id  []Tipo_contenido
	campos_dinamicos   []Campo_dinamico
	contenido          string
	estilos            []map[string]interface{}
	orden              int
	activo             bool
	fecha_creacion     string
	fecha_modificacion bool
}
