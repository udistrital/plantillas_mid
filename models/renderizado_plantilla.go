package models

type Renderizado_plantilla struct {
	Plantilla_id string                 `json:"plantilla_id"`
	Datos        map[string]interface{} `json:"data"`
}
