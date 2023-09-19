package models

import (
	"fmt"

	"github.com/udistrital/utils_oas/formatdata"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func RegistrarTitulo(titulo map[string]interface{}) (result map[string]interface{}, outputError interface{}) {

	var TituloPost map[string]interface{}
	var resultadoRegistro map[string]interface{}
	var errReg interface{}

	TituloPost = ConstruirTitulo(titulo)

	errReg = request.SendJson(beego.AppConfig.String("PlantillasCrudService")+"/titulo", "POST", &resultadoRegistro, TituloPost)

	if resultadoRegistro["Status"] == "400" || errReg != nil {
		fmt.Println(errReg)
		return nil, errReg

	} else {
		formatdata.JsonPrint(resultadoRegistro)
		return resultadoRegistro, nil

	}
}

func ConsultarTitulo(id string) (map[string]interface{}, error) {

	var titulo map[string]interface{}

	err1 := request.GetJson(beego.AppConfig.String("PlantillasCrudService")+id, &titulo)

	if err1 == nil {
		return titulo, nil
	} else {
		fmt.Println("Error al consultar la titulo: ", err1)
		return nil, err1
	}
}

func ConstruirTitulo(titulo map[string]interface{}) (tituloformatted map[string]interface{}) {
	Titulo := make(map[string]interface{})
	Titulo = titulo

	TituloPost := make(map[string]interface{})

	TituloPost = map[string]interface{}{
		"Activo":            true,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Nombre":            Titulo["nombre"],
		"Descripcion":       Titulo["descripcion"],
		"EnlaceDoc":         Titulo["enlaceDoc"],
		"Version":           Titulo["version"],
	}

	camposAdicionales := Titulo["camposAdicionales"].([]map[string]interface{})

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
	TituloPost["CamposAdicionales"] = campoAdicionalPost

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

	return TituloPost
}

func ActualizarTitulo(id string, titulo map[string]interface{}) (status interface{}, outputError interface{}) {

	var resultadoRegistro map[string]interface{}
	var err1 interface{}

	err1 = request.SendJson(beego.AppConfig.String("PlantillasCrudService")+"/titulo/"+id, "PUT", &resultadoRegistro, titulo)

	if err1 == nil {
		fmt.Println("Titulo actualizada!")
		return resultadoRegistro["res"], nil
	} else {
		return nil, err1
	}
}
