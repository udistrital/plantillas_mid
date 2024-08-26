package services

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/plantillas_mid/models"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
)

func ActualizarPlantilla(idStr string, registroPlantilla models.Plantilla) (interface{}, error) {
	defer errorhandler.HandlePanic(nil)

	result, err := ActualizarPlantilla(idStr, registroPlantilla)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func RegistrarPlantilla(plantilla models.Plantilla) (interface{}, error) {
	// Definir la URL del CRUD
	url := beego.AppConfig.String("PlantillasCrudService") + "/seccion"

	// Variable para almacenar el resultado del registro
	var resultadoRegistro interface{}

	// Enviar el JSON al CRUD utilizando la función SendJson
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

	// imagenesPost := make([]map[string]interface{}, 0)
	// for _, titulo := range titulos {
	// 	titulosPost = append(titulosPost, map[string]interface{}{
	// 		"Activo":            true,
	// 		"FechaCreacion":     nil,
	// 		"FechaModificacion": nil,
	// 		"Nombre":            titulo["nombre"],
	// 		"Descripcion":       titulo["descripcion"],
	// 		"Valor":             titulo["valor"],
	// 		"Posicion":          titulo["posicion"],
	// 		"CamposAdicionales": titulo["camposAdicionales"],
	// 		"EstilosFuente":     titulo["estilosFuente"],
	// 	})
	// }
	// PlantillaPost["Titulos"] = titulosPost

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
	// Implementar la lógica para obtener todas las plantillas
	return nil, nil
}
