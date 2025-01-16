package models

// Body que recibe la solicitud de /renderizar
type Renderizado struct {
	Css 		 string                  `json:"css"`
	Plantilla_id string                  `json:"plantilla_id"`
	Datos         map[string]interface{} `json:"datos"`
}

// Body que recibe la solicitud de /renderizarHTML
type RenderizadoHTML struct {
	Plantilla_id string                  `json:"plantilla_id"`
	Datos         map[string]interface{} `json:"datos"`
}

// Body que recibe la solicitud de /renderizarPDF
// Body que recibe la solicitud de /generar-pdf de PLANTILLAS_MID_RENDERIZADO_PLANTILLAS
type RenderizadoPDF struct {
	Css string 					  `json:"css"`
	Html string                   `json:"html"`
	Datos map[string]interface{}  `json:"datos"`
}

// Body que recibe la solicitud de /generar-html de PLANTILLAS_MID_RENDERIZADO_PLANTILLAS
type BodyHTML struct {
	Html string                  `json:"html"`
	Datos map[string]interface{} `json:"datos"`
}