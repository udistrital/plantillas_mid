package models

// Respuesta generada al procesar la solicitud /generar-html por PLANTILLAS_MID_RENDERIZADO_PLANTILLAS 
type ResponseHTML struct {
	Html    string `json:"html"`
	Message string `json:"Message"`
	Status  int    `json:"Status"`
	Success bool   `json:"Success"`
}

// Respuesta generada al procesar la solicitud /generar-pdf por PLANTILLAS_MID_RENDERIZADO_PLANTILLAS 
type ResponsePDF struct {
	Data    string `json:"Data"` // PDF en base64
	Message string `json:"Message"`
	Status  int    `json:"Status"`
	Success bool   `json:"Success"`
}
