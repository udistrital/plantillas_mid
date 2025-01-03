package models

// Body que recibe la solicitud de /renderizar
type Renderizado struct {
	Plantilla_id string                 `json:"plantilla_id"`
	Css 		 string                 `json:"css"`
	Data         map[string]interface{} `json:"data"`
}

// Body que recibe la solicitud de /renderizarHTML
type RenderizadoHTML struct {
	Plantilla_id string                 `json:"plantilla_id"`
	Data         map[string]interface{} `json:"data"`
}

// Body que recibe la solicitud de /renderizarPDF
// Body que recibe la solicitud de /generar-pdf de PLANTILLAS_MID_RENDERIZADO_PLANTILLAS
type RenderizadoPDF struct {
	Html string                  `json:"html"`
	Css string 					 `json:"css"`
	Data map[string]interface{}  `json:"data"`
}

// Body que recibe la solicitud de /generar-html de PLANTILLAS_MID_RENDERIZADO_PLANTILLAS
type BodyHTML struct {
	Html string                 `json:"html"`
	Data map[string]interface{} `json:"data"`
}