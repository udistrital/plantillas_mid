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

// Función auxiliar que convierte un objeto plantilla en HTML
// func ConversorHTML(plantilla Plantilla) (string, error) {
// 	var tmpl bytes.Buffer

// 	plantillaStruct := Plantilla{
// 		tipo_plantilla_id: plantilla.tipo_plantilla_id,
// 		nombre:            plantilla.nombre,
// 		secciones:         plantilla.secciones,
// 		fecha_creacion:    plantilla.fecha_creacion,
// 	}

// 	t := template.New("plantilla")
// 	_, err := t.Parse(plantillaStruct.tipo_plantilla_id)
// 	if err != nil {
// 		return "", err
// 	}

// 	err = t.Execute(&tmpl, plantillaStruct)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tmpl.String(), nil
// }

// Función auxiliar que convierte HTML a PDF usando gofpdf
