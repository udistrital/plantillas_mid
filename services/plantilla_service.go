package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/plantillas_mid/helpers"
	"github.com/udistrital/plantillas_mid/models"
)

type DocumentoResponse struct {
	Status string `json:"Status"`
	Res    struct {
		Enlace string `json:"Enlace"`
	} `json:"res"`
}

func DuplicarPlantilla(id string) (map[string]interface{}, error) {
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

	documentPayload := map[string]interface{}{
		"nombre":    plantillaDuplicada.Nombre,
		"contenido": plantillaDuplicada.Contenido,
	}

	documentPayloadBytes, err := json.Marshal(documentPayload)
	if err != nil {
		return nil, fmt.Errorf("error al convertir el payload del documento a JSON: %v", err)
	}

	req, err = http.NewRequest("POST", "https://autenticacion.portaloas.udistrital.edu.co/apioas/gestor_documental_mid/v1/document/store_document", bytes.NewBuffer(documentPayloadBytes))
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

	uid, ok := gestorDocumentalResponse["uid"].(string)
	if !ok {
		return nil, fmt.Errorf("el uid no se encuentra en la respuesta de gestor_documental_mid")
	}

	resultado["uid_documento"] = uid

	return resultado, nil
}

func ProcesarPlantilla(requestData map[string]interface{}) (string, error) {
	nombrePlantilla, ok := requestData["nombre_plantilla"].(string)
	if !ok {
		return "", errors.New("Falta 'nombre_plantilla' en el JSON")
	}

	htmlContent, ok := requestData["html"].(string)
	if !ok {
		return "", errors.New("Falta 'html' en el JSON")
	}

	base64PDF, err := helpers.ConvertHTMLToPDF(htmlContent, "", nombrePlantilla)
	if err != nil {
		return "", fmt.Errorf("error al convertir HTML a PDF: %v", err)
	}

	enlace, err := EnviarDocumento(base64PDF, nombrePlantilla)
	if err != nil {
		return "", fmt.Errorf("error al enviar documento: %v", err)
	}

	sistemaId, ok := requestData["sistema_id"].(string)
	if !ok {
		return "", errors.New("Falta 'sistema_id' en el JSON")
	}

	tipoPlantillaId, ok := requestData["tipo_plantilla_id"].(string)
	if !ok {
		return "", errors.New("Falta 'tipo_plantilla_id' en el JSON")
	}

	grupoId, ok := requestData["GrupoId"].(string)
	if !ok {
		return "", errors.New("Falta 'GrupoId' en el JSON")
	}

	plantilla := models.Plantilla{
		TipoPlantillaId: tipoPlantillaId,
		SistemaId:       sistemaId,
		Nombre:          nombrePlantilla,
		Contenido:       htmlContent,
		GrupoId:         grupoId,
		Activo:          true,
		Version:         1,
		Uid:             enlace,
	}

	err = models.AlmacenarPlantilla(plantilla)
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

	payload := models.CrearGestorDocumentalPayload(nombrePlantilla, base64PDF)

	var resultadoRegistro map[string]interface{}

	err = models.SendJson("http://"+gestorDocumentalURL+"/document/upload", "POST", &resultadoRegistro, payload)
	if err != nil {
		return "", fmt.Errorf("error al enviar documento: %v", err)
	}

	if resultadoRegistro["Status"].(string) == "200" {
		return resultadoRegistro["EnlaceDocumento"].(string), nil
	} else {
		return "", fmt.Errorf("error al registrar el documento: %v", resultadoRegistro["Error"].(string))
	}
}

func ProcesarDocumento(requestData map[string]interface{}) (string, error) {
	PlantillaId, ok := requestData["PlantillaId"].(string)
	if !ok {
		return "", errors.New("Falta 'PlantillaId' en el JSON")
	}

	nombrePlantilla, ok := requestData["nombre_plantilla"].(string)
	if !ok {
		return "", errors.New("Falta 'nombre_plantilla' en el JSON")
	}

	camposDinamicos, ok := requestData["campos_dinamicos"].(map[string]interface{})
	if !ok {
		return "", errors.New("Falta 'campos_dinamicos' en el JSON")
	}

	url := fmt.Sprintf("https://xxxxxxxxxxxxx/endpoint/%s", PlantillaId)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error al hacer la solicitud al endpoint: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error en la respuesta del endpoint: %s", resp.Status)
	}

	htmlContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error al leer la respuesta del endpoint: %v", err)
	}

	htmlProcesado := helpers.ReemplazarCamposDinamicos(string(htmlContent), camposDinamicos)

	base64PDF, err := helpers.ConvertHTMLToPDF(htmlProcesado, "", nombrePlantilla)
	if err != nil {
		return "", fmt.Errorf("error al convertir HTML a PDF: %v", err)
	}

	return base64PDF, nil
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
