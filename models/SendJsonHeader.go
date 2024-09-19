package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendJson(url string, method string, result interface{}, payload interface{}) error {
	// Serializar el payload a JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error al serializar el payload: %v", err)
	}

	// Crear la solicitud HTTP
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("error al crear la solicitud: %v", err)
	}

	// Definir el tipo de contenido como JSON
	req.Header.Set("Content-Type", "application/json")

	// Hacer la petición HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error en la petición HTTP: %v", err)
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error al leer la respuesta: %v", err)
	}

	// Verificar el estado de la respuesta HTTP
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error en la respuesta del servidor: %s", string(body))
	}

	// Deserializar la respuesta JSON
	err = json.Unmarshal(body, result)
	if err != nil {
		return fmt.Errorf("error al deserializar la respuesta: %v", err)
	}

	return nil
}
