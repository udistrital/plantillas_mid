package services

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/plantillas_mid/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
)

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

// CrearPlantilla - Crea una nueva plantilla en la base de datos.
func CrearPlantilla(plantilla models.Plantilla) (models.Plantilla, error) {
	var resultado models.Plantilla
	// Uso de request para insertar la plantilla en el backend correspondiente.
	if err := request.SendJson("POST", "/path/to/endpoint", &resultado, plantilla); err != nil {
		beego.Error("Error al crear plantilla: ", err)
		return models.Plantilla{}, err
	}
	return resultado, nil
}

// ActualizarPlantilla - Actualiza una plantilla existente por su ID.
func ActualizarPlantilla(id string, plantilla models.Plantilla) (models.Plantilla, error) {
	var resultado models.Plantilla
	url := fmt.Sprintf("/path/to/endpoint/%s", id)
	if err := request.SendJson("PUT", url, &resultado, plantilla); err != nil {
		beego.Error("Error al actualizar plantilla: ", err)
		return models.Plantilla{}, err
	}
	return resultado, nil
}

// DuplicarPlantilla - Duplicar una plantilla existente.
func DuplicarPlantilla(id string) (models.Plantilla, error) {
	var plantilla models.Plantilla
	url := fmt.Sprintf("/path/to/endpoint/%s", id)
	if err := request.GetJson(url, &plantilla); err != nil {
		beego.Error("Error al obtener plantilla para duplicar: ", err)
		return models.Plantilla{}, err
	}

	nuevaPlantilla := plantilla
	nuevaPlantilla.Id = "" // Nuevo ID para la plantilla duplicada

	var resultado models.Plantilla
	if err := request.SendJson("POST", "/path/to/endpoint", &resultado, nuevaPlantilla); err != nil {
		beego.Error("Error al duplicar plantilla: ", err)
		return models.Plantilla{}, err
	}
	return resultado, nil
}

// DeletePlantilla - Elimina una plantilla en función de su ID.
func DeletePlantilla(id string) error {
	url := fmt.Sprintf("/path/to/endpoint/%s", id)
	if err := request.SendJson("DELETE", url, nil, nil); err != nil {
		beego.Error("Error al eliminar plantilla: ", err)
		return err
	}
	return nil
}

// GetPlantillas - Obtiene todas las plantillas.
func GetPlantillas() ([]models.Plantilla, error) {
	var plantillas []models.Plantilla
	if err := request.GetJson("/path/to/endpoint", &plantillas); err != nil {
		beego.Error("Error al obtener plantillas: ", err)
		return nil, err
	}
	return plantillas, nil
}

// GetVersiones - Devuelve las versiones de una plantilla específica por su ID.
func GetVersiones(id string) ([]models.Plantilla, error) {
	var versiones []models.Plantilla
	url := fmt.Sprintf("/path/to/endpoint/versiones/%s", id)
	if err := request.GetJson(url, &versiones); err != nil {
		beego.Error("Error al obtener versiones de la plantilla: ", err)
		return nil, err
	}
	return versiones, nil
}
