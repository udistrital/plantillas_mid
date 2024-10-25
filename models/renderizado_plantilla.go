package models

type RenderizadoPlantilla struct {
	Plantilla_id string                 `json:"plantilla_id"`
	Datos        map[string]interface{} `json:"data"`
}
