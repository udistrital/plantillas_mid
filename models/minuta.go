package models

import (
	"fmt"

	"github.com/udistrital/utils_oas/formatdata"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func RegistrarMinuta(minuta map[string]interface{}) (result map[string]interface{}, outputError interface{}) {

	var MinutaPost map[string]interface{}
	var resultadoRegistro map[string]interface{}
	var errReg interface{}

	MinutaPost = ConstruirMinuta(minuta)

	errReg = request.SendJson(beego.AppConfig.String("PlantillasCrudService")+"/minuta", "POST", &resultadoRegistro, MinutaPost)

	if resultadoRegistro["Status"] == "400" || errReg != nil {
		fmt.Println(errReg)
		return nil, errReg

	} else {
		formatdata.JsonPrint(resultadoRegistro)
		return resultadoRegistro, nil

	}
}

func ConsultarMinuta(id string) (map[string]interface{}, error) {

	var minuta map[string]interface{}

	err1 := request.GetJson(beego.AppConfig.String("PlantillasCrudService")+id, &minuta)

	if err1 == nil {
		return minuta, nil
	} else {
		fmt.Println("Error al consultar la minuta: ", err1)
		return nil, err1
	}
}

func ConstruirMinuta(minuta map[string]interface{}) (minutaformatted map[string]interface{}) {
	Minuta := make(map[string]interface{})
	Minuta = minuta

	MinutaPost := make(map[string]interface{})

	MinutaPost = map[string]interface{}{
		"Activo":            true,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Nombre":            Minuta["nombre"],
		"Descripcion":       Minuta["descripcion"],
		"EnlaceDoc":         Minuta["enlaceDoc"],
		"Version":           Minuta["version"],
	}

	camposAdicionales := Minuta["camposAdicionales"].([]map[string]interface{})

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
	MinutaPost["CamposAdicionales"] = campoAdicionalPost

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

	return MinutaPost
}

func ActualizarMinuta(id string, minuta map[string]interface{}) (status interface{}, outputError interface{}) {

	var resultadoRegistro map[string]interface{}
	var err1 interface{}

	err1 = request.SendJson(beego.AppConfig.String("PlantillasCrudService")+"/minuta/"+id, "PUT", &resultadoRegistro, minuta)

	if err1 == nil {
		fmt.Println("Minuta actualizada!")
		return resultadoRegistro["res"], nil
	} else {
		return nil, err1
	}
}
