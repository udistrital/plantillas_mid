package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/udistrital/plantillas_mid/models"
)

func DuplicarPlantilla(id string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://mi-crud.com/plantillas/%s", id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("error al obtener la plantilla original del CRUD")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("error al obtener la plantilla original del CRUD: %s", string(body))
	}

	var plantillaOriginal models.Plantilla
	if err := json.NewDecoder(resp.Body).Decode(&plantillaOriginal); err != nil {
		return nil, errors.New("error al decodificar la plantilla original")
	}

	plantillaDuplicada := models.Plantilla{
		Id:                "",
		TipoPlantillaId:   plantillaOriginal.TipoPlantillaId,
		SistemaId:         plantillaOriginal.SistemaId,
		Nombre:            plantillaOriginal.Nombre + " - Copia",
		Contenido:         plantillaOriginal.Contenido,
		GrupoId:           plantillaOriginal.GrupoId,
		Version:           plantillaOriginal.Version + 1,
		Uid:               plantillaOriginal.Uid,
		Metadatos:         plantillaOriginal.Metadatos,
		Activo:            true,
		FechaCreacion:     time.Now(),
		FechaModificacion: time.Now(),
	}

	payload, err := json.Marshal(plantillaDuplicada)
	if err != nil {
		return nil, errors.New("error al convertir la plantilla duplicada a JSON")
	}

	req, err := http.NewRequest("POST", "https://mi-crud.com/plantillas", bytes.NewBuffer(payload))
	if err != nil {
		return nil, errors.New("error al crear la solicitud HTTP para duplicar la plantilla")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return nil, errors.New("error al enviar la solicitud al CRUD")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("error del CRUD al duplicar la plantilla: %s", string(body))
	}

	var resultado map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&resultado); err != nil {
		return nil, errors.New("error al decodificar la respuesta del CRUD")
	}

	return resultado, nil
}

func ConstruirPlantilla(body map[string]interface{}) (string, error) {
	contenido, ok := body["contenido"].(string)
	if !ok {
		return "", errors.New("el JSON no contiene un campo 'contenido' válido")
	}

	// Generar el PDF
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return "", errors.New("error al crear el generador de PDF")
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(contenido))))
	pdfg.MarginLeft.Set(10)
	pdfg.MarginRight.Set(10)
	pdfg.MarginTop.Set(10)
	pdfg.MarginBottom.Set(10)

	if err := pdfg.Create(); err != nil {
		return "", errors.New("error al generar el PDF")
	}

	pdfBytes := pdfg.Bytes()

	// Probar la generación del archivo guardándolo localmente
	if err := ioutil.WriteFile("plantilla_generada.pdf", pdfBytes, 0644); err != nil {
		return "", errors.New("error al guardar el PDF localmente para pruebas")
	}

	// Supuesto almacenamiento en Nuxeo
	uid, err := almacenarEnNuxeo(pdfBytes)
	if err != nil {
		return "", err
	}

	return uid, nil
}

// Función que simula el almacenamiento en Nuxeo
func almacenarEnNuxeo(pdfBytes []byte) (string, error) {
	// Aquí agregarías la lógica real para subir el PDF a Nuxeo.
	// Este es un ejemplo hipotético de cómo podrías hacerlo.

	url := "https://nuxeo-server.com/api/v1/upload" // Suponiendo que esta es la URL de Nuxeo
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(pdfBytes))
	if err != nil {
		return "", errors.New("error al crear la solicitud HTTP para Nuxeo")
	}

	req.Header.Set("Content-Type", "application/pdf")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("error al enviar la solicitud a Nuxeo")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("error del servidor Nuxeo: %s", string(body))
	}

	// Supongamos que Nuxeo devuelve el UID en el cuerpo de la respuesta
	var nuxeoResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&nuxeoResponse); err != nil {
		return "", errors.New("error al decodificar la respuesta de Nuxeo")
	}

	uid, ok := nuxeoResponse["uid"].(string)
	if !ok {
		return "", errors.New("la respuesta de Nuxeo no contiene un UID válido")
	}

	return uid, nil
}
