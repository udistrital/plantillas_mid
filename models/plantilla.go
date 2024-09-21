package models

import (
	"time"
)

type Plantilla struct {
	Id                string                 `json:"_id" bson:"_id"`
	TipoPlantillaId   string                 `json:"tipo_plantilla_id" bson:"tipo_plantilla_id"`
	SistemaId         string                 `json:"sistema_id" bson:"sistema_id"`
	Nombre            string                 `json:"nombre" bson:"nombre"`
	Contenido         string                 `json:"contenido" bson:"contenido"`
	GrupoId           string                 `json:"grupo_id" bson:"grupo_id"`
	Version           int                    `json:"version" bson:"version"`
	Uid               string                 `json:"uid" bson:"uid"`
	Metadatos         map[string]interface{} `json:"metadatos" bson:"metadatos"`
	Activo            bool                   `json:"activo" bson:"activo"`
	FechaCreacion     time.Time              `json:"fecha_creacion" bson:"fecha_creacion"`
	FechaModificacion time.Time              `json:"fecha_modificacion" bson:"fecha_modificacion"`
}
