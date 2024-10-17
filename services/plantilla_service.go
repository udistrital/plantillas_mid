package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plantillas_mid/models"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
)

func ActualizarPlantilla(idStr string, registroPlantilla models.Plantilla) (interface{}, error) {
	defer errorhandler.HandlePanic(nil)

	result, err := models.ActualizarPlantilla(idStr, registroPlantilla)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func RegistrarPlantilla(plantilla models.Plantilla) (interface{}, error) {
	url := beego.AppConfig.String("PlantillasCrudService") + "/seccion"

	var resultadoRegistro interface{}

	err := request.SendJson(url, "POST", &resultadoRegistro, plantilla)
	if err != nil {
		return nil, errors.New("error al registrar la plantilla en el CRUD: " + err.Error())
	}

	return resultadoRegistro, nil
}

func ConstruirSeccion(seccion map[string]interface{}) (seccionformatted map[string]interface{}) {
	Seccion := make(map[string]interface{})
	Seccion = seccion

	SeccionPost := make(map[string]interface{})

	SeccionPost = map[string]interface{}{
		"Activo":            true,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Nombre":            Seccion["nombre"],
		"Descripcion":       Seccion["descripcion"],
		"EnlaceDoc":         Seccion["enlaceDoc"],
		"Version":           Seccion["version"],
	}

	camposAdicionales := Seccion["camposAdicionales"].([]map[string]interface{})

	campoAdicionalPost := make([]map[string]interface{}, 0)
	for _, campo := range camposAdicionales {
		campoAdicionalPost = append(campoAdicionalPost, map[string]interface{}{
			"Activo":            true,
			"FechaCreacion":     nil,
			"FechaModificacion": nil,
			"Nombre":            campo["nombre"],
			"Descripcion":       campo["descripcion"],
			"Valor":             campo["valor"],
			"Posicion":          campo["posicion"],
			"EstilosFuente":     campo["estilosFuente"],
		})
	}
	SeccionPost["CamposAdicionales"] = campoAdicionalPost

	return SeccionPost
}

func RegistrarSeccion(seccion map[string]interface{}) (result map[string]interface{}, outputError interface{}) {

	var SeccionPost map[string]interface{}
	var resultadoRegistro map[string]interface{}
	var errReg interface{}

	SeccionPost = ConstruirSeccion(seccion)

	errReg = request.SendJson(beego.AppConfig.String("PlantillasCrudService")+"/seccion", "POST", &resultadoRegistro, SeccionPost)

	if resultadoRegistro["Status"] == "400" || errReg != nil {
		fmt.Println(errReg)
		return nil, errReg

	} else {
		formatdata.JsonPrint(resultadoRegistro)
		return resultadoRegistro, nil

	}
}

func EditPlantillas() (interface{}, error) {
	return nil, nil
}

func DefinirPlantilla(requestData models.Renderizado_plantilla) (string, error) {
	logs.Info(requestData.Plantilla_id)
	html, err := TraerPlantilla(requestData.Plantilla_id)
	//logs.Info("HTML: ", html)
	if err != nil {
		fmt.Println("err1")
		logs.Error(err)
		return "", err
	}
	plantillaRenderizada, err := Renderizar(html, requestData.Datos)
	if err != nil {
		fmt.Println("err2")
		logs.Error(err)
		return "", err
	}
	logs.Info(plantillaRenderizada)
	return plantillaRenderizada, nil
}

func TraerPlantilla(plantilla_id string) (*string, error) {
	var response map[string]interface{}
	apiUrl := beego.AppConfig.String("plantillasServerless")

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
