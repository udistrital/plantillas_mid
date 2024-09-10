package services

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/udistrital/plantillas_mid/models"
)

func DuplicarPlantilla(id string) (map[string]interface{}, error) {
	// URL del CRUD para obtener la plantilla original
	url := fmt.Sprintf("https://2hzmbx74aj.execute-api.us-east-1.amazonaws.com/Prod/plantillas/%s", id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error al obtener la plantilla original del CRUD: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("error al obtener la plantilla original del CRUD: %s", string(body))
	}

	var plantillaOriginal models.Plantilla
	if err := json.NewDecoder(resp.Body).Decode(&plantillaOriginal); err != nil {
		return nil, fmt.Errorf("error al decodificar la plantilla original: %v", err)
	}

	// Crear la plantilla duplicada eliminando campos generados automáticamente
	plantillaDuplicada := models.Plantilla{
		TipoPlantillaId: plantillaOriginal.TipoPlantillaId,
		SistemaId:       plantillaOriginal.SistemaId,
		Nombre:          plantillaOriginal.Nombre + " - Copia",
		Contenido:       plantillaOriginal.Contenido,
		GrupoId:         plantillaOriginal.GrupoId,
		Version:         plantillaOriginal.Version + 1,
		Uid:             plantillaOriginal.Uid,
		Metadatos:       plantillaOriginal.Metadatos,
		Activo:          true,
	}

	payload, err := json.Marshal(plantillaDuplicada)
	if err != nil {
		return nil, fmt.Errorf("error al convertir la plantilla duplicada a JSON: %v", err)
	}

	// URL del CRUD para crear la nueva plantilla
	req, err := http.NewRequest("POST", "https://2hzmbx74aj.execute-api.us-east-1.amazonaws.com/Prod/plantillas", bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("error al crear la solicitud HTTP para duplicar la plantilla: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error al enviar la solicitud al CRUD: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("error del CRUD al duplicar la plantilla: %s", string(body))
	}

	var resultado map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&resultado); err != nil {
		return nil, fmt.Errorf("error al decodificar la respuesta del CRUD: %v", err)
	}

	return resultado, nil
}

type PDFRequest struct {
	HTML  string                 `json:"html"`
	CSS   string                 `json:"css,omitempty"`
	Datos map[string]interface{} `json:"datos,omitempty"`
}

type PDFResponse struct {
	Message string `json:"Message"`
	Success bool   `json:"Success"`
	Status  int    `json:"Status"`
	Data    string `json:"Data"`
}

func RenderizarPDF(html, css string, datos map[string]interface{}) ([]byte, error) {
	url := "http://localhost:RENDERIZADO_HTML_PORT/pdf"

	pdfRequest := PDFRequest{
		HTML:  html,
		CSS:   css,
		Datos: datos,
	}

	jsonData, err := json.Marshal(pdfRequest)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pdfResponse PDFResponse
	err = json.NewDecoder(resp.Body).Decode(&pdfResponse)
	if err != nil {
		return nil, err
	}

	if !pdfResponse.Success {
		return nil, errors.New(pdfResponse.Message)
	}

	pdfData, err := base64.StdEncoding.DecodeString(pdfResponse.Data)
	if err != nil {
		return nil, err
	}

	return pdfData, nil
}

func ComprobarConexionCRUD() error {
	url := "https://2hzmbx74aj.execute-api.us-east-1.amazonaws.com/Prod/health"

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("error al intentar conectar con el CRUD: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("el CRUD no está disponible o respondió con un código de estado no esperado")
	}

	return nil
}
