package models

// Respuesta obtenida al procesar la solicitud /generar-pdf o /generar-html de PLANTILLAS_MID_RENDERIZADO_PLANTILLAS
type ResponseRenderizado struct {
	Data    string `json:"Data"`
	Message string `json:"Message"`
	Status  int    `json:"Status"`
	Success bool   `json:"Success"`
}
