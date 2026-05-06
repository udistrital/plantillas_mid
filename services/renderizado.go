package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plantillas_mid/models"
	"github.com/udistrital/utils_oas/errorhandler"
)

func Renderizar(request models.Renderizado) (string, error) {

	defer errorhandler.HandlePanic(nil)

	html, err := ObtenerPlantilla(request.Plantilla_id)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	requestBody := models.RenderizadoPDF{
		Html: *html,
		Css:  request.Css,
		Data: request.Data,
	}

	plantillaRenderizada, err := GenerarPDF(requestBody)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	return plantillaRenderizada, nil
}

func RenderizarPDF(request models.RenderizadoPDF) (string, error) {
	defer errorhandler.HandlePanic(nil)

	plantillaRenderizada, err := GenerarPDF(request)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	return plantillaRenderizada, nil
}

func RenderizarDOCX(request models.Renderizado) (string, error) {
	defer errorhandler.HandlePanic(nil)

	html, err := ObtenerPlantilla(request.Plantilla_id)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	requestBody := models.RenderizadoDOCX{
		Html: *html,
		Css:  request.Css,
		Data: request.Data,
	}

	plantillaRenderizada, err := GenerarDOCX(requestBody)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	return plantillaRenderizada, nil
}

func GenerarPDF(requestBody models.RenderizadoPDF) (string, error) {
	defer errorhandler.HandlePanic(nil)

	response := models.ResponseRenderizado{}
	err := Generar(requestBody, &response, "generar-pdf")
	if err != nil {
		return "", err
	}
	return response.Data, nil
}

func GenerarDOCX(requestBody models.RenderizadoDOCX) (string, error) {
	defer errorhandler.HandlePanic(nil)

	response := models.ResponseRenderizado{}
	err := Generar(requestBody, &response, "generar-docx")
	if err != nil {
		return "", err
	}
	return response.Data, nil
}

func RenderizarHTML(request models.RenderizadoHTML) (string, error) {
	defer errorhandler.HandlePanic(nil)

	html, err := ObtenerPlantilla(request.Plantilla_id)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	plantillaRenderizada, err := GenerarHTML(html, request.Data)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	return plantillaRenderizada, nil
}

func GenerarHTML(html *string, datos map[string]interface{}) (string, error) {
	defer errorhandler.HandlePanic(nil)

	requestBody := models.BodyHTML{
		Html: *html,
		Data: datos,
	}

	response := models.ResponseRenderizado{}
	err := Generar(requestBody, &response, "generar-html")
	if err != nil {
		return "", err
	}
	return response.Data, nil
}

func ObtenerPlantilla(plantilla_id string) (*string, error) {
	defer errorhandler.HandlePanic(nil)

	apiUrl := beego.AppConfig.String("PlantillasCrudService")
	url := fmt.Sprintf("%s"+"/plantilla/"+"%s", apiUrl, plantilla_id)
	logs.Info("--------url OBTENER PLANTILLA: ", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error al obtener la plantilla: %s", string(body))
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	data, ok := response["Data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("error: no se pudo obtener 'Data' de la respuesta")
	}

	contenido, ok := data["contenido"].(string)
	if !ok {
		return nil, fmt.Errorf("error: 'contenido' no es de tipo string o no existe en 'Data'")
	}
	return &contenido, nil
}

func Generar(requestBody interface{}, responseBody interface{}, endpoint string) error {
	defer errorhandler.HandlePanic(nil)

	// Serializar el cuerpo de la solicitud
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	// Construir la URL
	api := beego.AppConfig.String("RenderizadoPlantillas")
	url := fmt.Sprintf("%s/%s", api, endpoint)
	logs.Info("--------url RENDERIZADO: ", url)

	// Crear la solicitud HTTP
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Leer el cuerpo de la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Deserializar la respuesta en el tipo proporcionado
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return err
	}
	return nil
}
