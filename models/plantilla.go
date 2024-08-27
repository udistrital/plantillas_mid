package models

import (
	"time"
)

type Plantilla struct {
	id                 string
	tipo_plantilla_id  string
	secciones          []Seccion
	sistema_id         string
	grupo_id           string
	version            int
	nombre             string
	uid                bool
	activo             bool
	fecha_creacion     time.Time
	fecha_modificacion time.Time
}

func ActualizarPlantilla(idStr string, registroPlantilla Plantilla) (interface{}, error) {
	return nil, nil
}

func EliminarPlantilla(idStr string) error {
	return nil
}
