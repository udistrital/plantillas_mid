package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/astaxie/beego"
)

type Plantilla struct {
	Id                string                 `json:"_id" bson:"_id"`
	TipoPlantillaId   string                 `json:"tipo_plantilla_id" bson:"tipo_plantilla_id"`
	SistemaId         string                 `json:"sistema_id" bson:"sistema_id"`
	Nombre            string                 `json:"nombre" bson:"nombre"`
	Contenido         string                 `json:"contenido" bson:"contenido"`
	GrupoId           string                 `json:"grupo_id" bson:"grupo_id"`
	Version           int                    `json:"version" bson:"version"`
	Uid               string                 `json:"uid" bson:"uid"`
	Metadatos         map[string]interface{} `json:"metadatos" bson:"metadatos"`
	Activo            bool                   `json:"activo" bson:"activo"`
	FechaCreacion     time.Time              `json:"fecha_creacion" bson:"fecha_creacion"`
	FechaModificacion time.Time              `json:"fecha_modificacion" bson:"fecha_modificacion"`
}

func AlmacenarPlantilla(plantilla Plantilla) error {
	crudServiceURL := beego.AppConfig.String("PlantillasCrudService")
	if crudServiceURL == "" {
		return fmt.Errorf("la variable de entorno 'PlantillasCrudService' no está definida")
	}

	payloadBytes, err := json.Marshal(plantilla)
	if err != nil {
		return fmt.Errorf("error al serializar la plantilla: %v", err)
	}

	resp, err := http.Post("http://"+crudServiceURL+"/plantillas", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("error en la petición POST al CRUD de plantillas: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("error en la respuesta del CRUD de plantillas: %s", string(bodyBytes))
	}

	return nil
}
