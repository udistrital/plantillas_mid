package models

import (
	"fmt"

	"github.com/udistrital/utils_oas/formatdata"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

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

func ConsultarSeccion(id string) (map[string]interface{}, error) {

	var seccion map[string]interface{}

	err1 := request.GetJson(beego.AppConfig.String("PlantillasCrudService")+id, &seccion)

	if err1 == nil {
		return seccion, nil
	} else {
		fmt.Println("Error al consultar la seccion: ", err1)
		return nil, err1
	}
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

func ActualizarSeccion(id string, seccion map[string]interface{}) (status interface{}, outputError interface{}) {

	var resultadoRegistro map[string]interface{}
	var err1 interface{}

	err1 = request.SendJson(beego.AppConfig.String("PlantillasCrudService")+"/seccion/"+id, "PUT", &resultadoRegistro, seccion)

	if err1 == nil {
		fmt.Println("Seccion actualizada!")
		return resultadoRegistro["res"], nil
	} else {
		return nil, err1
	}
}
