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

func DefinirPlantilla(requestData models.RenderizadoPlantilla) (string, error) {
	defer errorhandler.HandlePanic(nil)

	html, err := TraerPlantilla(requestData.Plantilla_id)
	//logs.Info("HTML: ", html)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	plantillaRenderizada, err := Renderizar(html, requestData.Datos)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	return plantillaRenderizada, nil
}

func TraerPlantilla(plantilla_id string) (*string, error) {
	defer errorhandler.HandlePanic(nil)

	var response map[string]interface{}
	apiUrl := beego.AppConfig.String("PlantillasCrudService")

	url := fmt.Sprintf("%s"+"/plantilla/"+"%s", apiUrl, plantilla_id)
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

	//logs.Info("Contenido de la plantilla: ", contenido)
	return &contenido, nil
}

func Renderizar(html *string, data map[string]interface{}) (string, error) {
	defer errorhandler.HandlePanic(nil)

	var response models.Response
	url := beego.AppConfig.String("RenderizadoPlantillas")

	requestBody := map[string]interface{}{
		"html": *html,
		"data": data,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}
	return response.Data, nil
}
