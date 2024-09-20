package helpers

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type GestorDocumentalPayload struct {
	IdTipoDocumento int               `json:"IdTipoDocumento"`
	Nombre          string            `json:"nombre"`
	Metadatos       map[string]string `json:"metadatos"`
	Descripcion     string            `json:"descripcion"`
	File            string            `json:"file"`
}
type PlantillaParams struct {
	NombrePlantilla string
	HTMLContent     string
	SistemaId       string
	TipoPlantillaId string
	GrupoId         string
}

func ConvertHTMLToPDF(htmlContent, css string, outputFilename string) (string, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Printf("Error al inicializar el generador PDF: %v", err)
		return "", fmt.Errorf("error al inicializar el generador PDF: %v", err)
	}

	page := wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(htmlContent)))

	if css != "" {
		page.EnableLocalFileAccess.Set(true)
		page.UserStyleSheet.Set(css)
	}

	pdfg.AddPage(page)

	pdfg.MarginTop.Set(10)
	pdfg.MarginBottom.Set(10)
	pdfg.MarginLeft.Set(10)
	pdfg.MarginRight.Set(10)

	err = pdfg.Create()
	if err != nil {
		log.Printf("Error al crear el PDF: %v", err)
		return "", fmt.Errorf("error al crear el PDF: %v", err)
	}

	pdfBuffer := pdfg.Bytes()

	base64PDF := base64.StdEncoding.EncodeToString(pdfBuffer)

	return base64PDF, nil
}

func ReemplazarCamposDinamicos(htmlContent string, dynamicFields map[string]interface{}) string {
	for key, value := range dynamicFields {
		placeholder := fmt.Sprintf("{{%s}}", key)

		log.Printf("Reemplazando placeholder '%s' con el valor '%v'", placeholder, value)

		htmlContent = strings.ReplaceAll(htmlContent, placeholder, fmt.Sprintf("%v", value))
	}

	return htmlContent
}

func CrearGestorDocumentalPayload(nombrePlantilla, base64PDF string) []GestorDocumentalPayload {

	return []GestorDocumentalPayload{
		{
			IdTipoDocumento: 170,
			Nombre:          nombrePlantilla,
			Metadatos:       map[string]string{},
			Descripcion:     "Plantillas",
			File:            base64PDF,
		},
	}
}

func ValidarPlantillaParametros(requestData map[string]interface{}) (*PlantillaParams, error) {
	params := &PlantillaParams{}

	if nombrePlantilla, ok := requestData["nombre_plantilla"].(string); ok {
		params.NombrePlantilla = nombrePlantilla
	} else {
		return nil, errors.New("Falta 'nombre_plantilla' en el JSON")
	}

	if htmlContent, ok := requestData["html"].(string); ok {
		params.HTMLContent = htmlContent
	} else {
		return nil, errors.New("Falta 'html' en el JSON")
	}

	if sistemaId, ok := requestData["sistema_id"].(string); ok {
		params.SistemaId = sistemaId
	} else {
		return nil, errors.New("Falta 'sistema_id' en el JSON")
	}

	if tipoPlantillaId, ok := requestData["tipo_plantilla_id"].(string); ok {
		params.TipoPlantillaId = tipoPlantillaId
	} else {
		return nil, errors.New("Falta 'tipo_plantilla_id' en el JSON")
	}

	if grupoId, ok := requestData["GrupoId"].(string); ok {
		params.GrupoId = grupoId
	} else {
		return nil, errors.New("Falta 'GrupoId' en el JSON")
	}

	return params, nil
}

type DocumentoParams struct {
	PlantillaId     string
	NombrePlantilla string
	CamposDinamicos map[string]interface{}
}

func ValidarDocumentoParametros(requestData map[string]interface{}) (*DocumentoParams, error) {
	params := &DocumentoParams{}

	if plantillaId, ok := requestData["PlantillaId"].(string); ok {
		params.PlantillaId = plantillaId
	} else {
		return nil, errors.New("Falta 'PlantillaId' en el JSON")
	}

	if nombrePlantilla, ok := requestData["nombre_plantilla"].(string); ok {
		params.NombrePlantilla = nombrePlantilla
	} else {
		return nil, errors.New("Falta 'nombre_plantilla' en el JSON")
	}

	if camposDinamicos, ok := requestData["campos_dinamicos"].(map[string]interface{}); ok {
		params.CamposDinamicos = camposDinamicos
	} else {
		return nil, errors.New("Falta 'campos_dinamicos' en el JSON")
	}

	return params, nil
}
