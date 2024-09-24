package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/udistrital/plantillas_mid/helpers"
	"github.com/udistrital/plantillas_mid/models"
	"github.com/udistrital/utils_oas/request"
)

type DocumentoResponse struct {
	Status string `json:"Status"`
	Res    struct {
		Enlace string `json:"Enlace"`
	} `json:"res"`
}

func DuplicarPlantilla(id string) (map[string]interface{}, error) {

	plantillasCrudService := beego.AppConfig.String("PlantillasCrudService")
	if plantillasCrudService == "" {
		return nil, errors.New("la variable de entorno 'PlantillasCrudService' no está configurada")
	}

	url := fmt.Sprintf("http://%s/plantilla/%s", plantillasCrudService, id)

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

	plantillaDuplicada := models.Plantilla{
		TipoPlantillaId: plantillaOriginal.TipoPlantillaId,
		SistemaId:       plantillaOriginal.SistemaId,
		Nombre:          plantillaOriginal.Nombre + " (versión duplicada)",
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

	req, err := http.NewRequest("POST", "http://"+plantillasCrudService+"/plantilla", bytes.NewBuffer(payload))
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

	gestorDocumentalService := beego.AppConfig.String("GestorDocumental")
	if gestorDocumentalService == "" {
		return nil, errors.New("la variable de entorno 'GestorDocumental' no está configurada")
	}

	documentPayload := map[string]interface{}{
		"nombre":    plantillaDuplicada.Nombre,
		"contenido": plantillaDuplicada.Contenido,
	}

	documentPayloadBytes, err := json.Marshal(documentPayload)
	if err != nil {
		return nil, fmt.Errorf("error al convertir el payload del documento a JSON: %v", err)
	}

	req, err = http.NewRequest("POST", "http://"+gestorDocumentalService+"/document/store_document", bytes.NewBuffer(documentPayloadBytes))
	if err != nil {
		return nil, fmt.Errorf("error al crear la solicitud HTTP para almacenar el documento en gestor_documental_mid: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error al enviar la solicitud a gestor_documental_mid: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("error al almacenar el documento en gestor_documental_mid: %s", string(body))
	}

	var gestorDocumentalResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&gestorDocumentalResponse); err != nil {
		return nil, fmt.Errorf("error al decodificar la respuesta de gestor_documental_mid: %v", err)
	}

	uid, ok := gestorDocumentalResponse["Enlace"].(string)
	if !ok {
		return nil, fmt.Errorf("el uid no se encuentra en la respuesta de gestor_documental_mid")
	}

	resultado["uid_documento"] = uid

	return resultado, nil
}

func ProcesarPlantilla(requestData map[string]interface{}) (string, error) {
	params, err := helpers.ValidarPlantillaParametros(requestData)
	if err != nil {
		return "", err
	}

	base64PDF, err := helpers.ConvertHTMLToPDF(params.HTMLContent, "", params.NombrePlantilla)
	if err != nil {
		return "", fmt.Errorf("error al convertir HTML a PDF: %v", err)
	}

	enlace, err := EnviarDocumento(base64PDF, params.NombrePlantilla)
	if err != nil {
		return "", fmt.Errorf("error al enviar documento: %v", err)
	}

	plantilla := models.Plantilla{
		TipoPlantillaId: params.TipoPlantillaId,
		SistemaId:       params.SistemaId,
		Nombre:          params.NombrePlantilla,
		Contenido:       params.HTMLContent,
		GrupoId:         params.GrupoId,
		Activo:          true,
		Version:         1,
		Uid:             enlace,
	}

	err = helpers.AlmacenarPlantilla(plantilla)
	if err != nil {
		return "", fmt.Errorf("error al almacenar la plantilla: %v", err)
	}

	return enlace, nil
}

func EnviarDocumento(base64PDF, nombrePlantilla string) (enlace string, err error) {
	gestorDocumentalURL := beego.AppConfig.String("GestorDocumental")
	if gestorDocumentalURL == "" {
		return "", fmt.Errorf("la variable de entorno 'GestorDocumental' no está definida")
	}

	payload := helpers.CrearGestorDocumentalPayload(nombrePlantilla, base64PDF)

	var resultadoRegistro map[string]interface{}

	err = request.SendJson("http://"+gestorDocumentalURL+"/document/store_document", "POST", &resultadoRegistro, payload)
	if err != nil {
		return "", fmt.Errorf("error al enviar documento: %v", err)
	}

	if resultadoRegistro["Status"].(string) == "200" {
		return resultadoRegistro["Enlace"].(string), nil
	} else {
		return "", fmt.Errorf("error al registrar el documento: %v", resultadoRegistro["Error"].(string))
	}
}

func ProcesarDocumento(requestData map[string]interface{}) (string, error) {
	params, err := helpers.ValidarDocumentoParametros(requestData)
	if err != nil {
		return "", err
	}

	consulta := fmt.Sprintf("%s/plantilla/%s", beego.AppConfig.String("PlantillasCrudService"), params.PlantillaId)

	var plantillaResponse map[string]interface{}
	if err := request.GetJson(consulta, &plantillaResponse); err != nil {
		return "", fmt.Errorf("error al hacer la solicitud al servicio de plantillas: %v", err)
	}

	htmlContent, ok := plantillaResponse["html_content"].(string)
	if !ok {
		return "", errors.New("El contenido HTML no se encuentra en la respuesta del servicio")
	}

	htmlProcesado := helpers.ReemplazarCamposDinamicos(htmlContent, params.CamposDinamicos)

	base64PDF, err := helpers.ConvertHTMLToPDF(htmlProcesado, "", params.NombrePlantilla)
	if err != nil {
		return "", fmt.Errorf("error al convertir HTML a PDF: %v", err)
	}

	return base64PDF, nil
}
