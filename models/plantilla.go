package models

import (
	"fmt"

	"github.com/udistrital/utils_oas/formatdata"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

func RegistrarPlantilla(plantilla map[string]interface{}) (result map[string]interface{}, outputError interface{}) {

	var PlantillaPost map[string]interface{}
	var resultadoRegistro map[string]interface{}
	var errRegPlantilla interface{}

	PlantillaPost = ConstruirPlantilla(plantilla)

	errRegPlantilla = request.SendJson(beego.AppConfig.String("PlantillasCrudService")+"/plantilla", "POST", &resultadoRegistro, PlantillaPost)

	if resultadoRegistro["Status"] == "400" || errRegPlantilla != nil {
		fmt.Println(errRegPlantilla)
		return nil, errRegPlantilla

	} else {
		formatdata.JsonPrint(resultadoRegistro)
		return resultadoRegistro, nil

	}
}

func ConsultarPlantilla() (plantilla map[string]interface{}, err error) {
	return nil, nil
}

func ConstruirPlantilla(plantilla map[string]interface{}) (plantillaFormatted map[string]interface{}) {
	Plantilla := make(map[string]interface{})
	Plantilla = plantilla

	PlantillaPost := make(map[string]interface{})

	secciones := Plantilla["secciones"].([]map[string]interface{})
	minutas := Plantilla["minutas"].([]map[string]interface{})
	titulos := Plantilla["titulos"].([]map[string]interface{})
	// imagenes := Plantilla["imagenes"].([]map[string]interface{})

	seccionesPost := make([]map[string]interface{}, 0)
	for _, seccion := range secciones {
		seccionesPost = append(seccionesPost, map[string]interface{}{
			"Activo":            true,
			"FechaCreacion":     nil,
			"FechaModificacion": nil,
			"Nombre":            seccion["nombre"],
			"Descripcion":       seccion["descripcion"],
			"Valor":             seccion["valor"],
			"Posicion":          seccion["posicion"],
			"CamposAdicionales": seccion["camposAdicionales"],
			"EstilosFuente":     seccion["estilosFuente"],
		})
	}
	PlantillaPost["Secciones"] = seccionesPost

	minutasPost := make([]map[string]interface{}, 0)
	for _, minuta := range minutas {
		minutasPost = append(minutasPost, map[string]interface{}{
			"Activo":            true,
			"FechaCreacion":     nil,
			"FechaModificacion": nil,
			"Nombre":            minuta["nombre"],
			"Descripcion":       minuta["descripcion"],
			"Valor":             minuta["valor"],
			"Posicion":          minuta["posicion"],
			"CamposAdicionales": minuta["camposAdicionales"],
			"EstilosFuente":     minuta["estilosFuente"],
		})
	}
	PlantillaPost["Minutas"] = minutasPost

	titulosPost := make([]map[string]interface{}, 0)
	for _, titulo := range titulos {
		titulosPost = append(titulosPost, map[string]interface{}{
			"Activo":            true,
			"FechaCreacion":     nil,
			"FechaModificacion": nil,
			"Nombre":            titulo["nombre"],
			"Descripcion":       titulo["descripcion"],
			"Valor":             titulo["valor"],
			"Posicion":          titulo["posicion"],
			"CamposAdicionales": titulo["camposAdicionales"],
			"EstilosFuente":     titulo["estilosFuente"],
		})
	}
	PlantillaPost["Titulos"] = titulosPost

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

	PlantillaPost = map[string]interface{}{
		"Activo":            true,
		"FechaCreacion":     nil,
		"FechaModificacion": nil,
		"Nombre":            Plantilla["nombre"],
		"Descripcion":       Plantilla["descripcion"],
		"EnlaceDoc":         Plantilla["enlaceDoc"],
		"Version":           Plantilla["version"],
	}

	return PlantillaPost
}

func ActualizarPlantilla(id string) (status interface{}, outputError interface{}) {
	return nil, nil
}
