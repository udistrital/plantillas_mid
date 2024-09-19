package models

type GestorDocumentalPayload struct {
	IdTipoDocumento int               `json:"IdTipoDocumento"`
	Nombre          string            `json:"nombre"`
	Metadatos       map[string]string `json:"metadatos"`
	Descripcion     string            `json:"descripcion"`
	File            string            `json:"file"`
}

func CrearGestorDocumentalPayload(nombrePlantilla, base64PDF string) []GestorDocumentalPayload {
	return []GestorDocumentalPayload{
		{
			IdTipoDocumento: 170,
			Nombre:          nombrePlantilla,
			Metadatos: map[string]string{
				"dato_a": "string",
				"dato_b": "string",
				"dato_n": "string",
			},
			Descripcion: "Plantillas",
			File:        base64PDF,
		},
	}
}
