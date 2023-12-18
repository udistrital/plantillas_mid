package models

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
)

type Plantilla struct {
	Id                int
	Tipo              string
	Nombre            string
	Descripcion       string
	Contenido         string
	EnlaceDoc         string
	FechaCreacion     time.Time
	FechaModificacion time.Time
	version           float64
	versionActual     bool
}

func RegistrarPlantilla(plantilla map[string]interface{}) (io.Reader, interface{}) {

	var PlantillaPost map[string]interface{}
	var resultadoRegistro map[string]interface{}
	var errRegPlantilla interface{}

	dataPlantilla := plantilla["contenido"].(string)

	pdfResponse, errConvert := ConvertHTMLToPDF(dataPlantilla)
	if errConvert != nil {
		fmt.Println("Error al convertir el html a pdf: ", errConvert)
		return nil, errConvert
	} else {
		fmt.Println("Html convertido a pdf correctamente", pdfResponse)
		PlantillaPost = plantilla

		fmt.Println("PlantillaPost: ", PlantillaPost)
		errRegPlantilla = request.SendJson(beego.AppConfig.String("PlantillasCrudService")+"/plantilla", "POST", &resultadoRegistro, PlantillaPost)

		if resultadoRegistro["Status"] == "400" || errRegPlantilla != nil {
			fmt.Println(errRegPlantilla)
			return nil, errRegPlantilla

		} else {
			formatdata.JsonPrint(resultadoRegistro)
			return nil, nil

		}
	}
}

func ConsultarPlantilla(id string) (map[string]interface{}, error) {

	var plantilla map[string]interface{}

	err1 := request.GetJson(beego.AppConfig.String("PlantillasCrudService")+id, &plantilla)

	if err1 == nil {
		return plantilla, nil
	} else {
		fmt.Println("Error al consultar la plantilla: ", err1)
		return nil, err1
	}
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

func ActualizarPlantilla(id string, plantilla map[string]interface{}) (status interface{}, outputError interface{}) {

	var resultadoRegistro map[string]interface{}
	var err1 interface{}

	err1 = request.SendJson(beego.AppConfig.String("PlantillasCrudService")+"/plantilla/"+id, "PUT", &resultadoRegistro, plantilla)

	if err1 == nil {
		fmt.Println("Plantilla actualizada!")
		return resultadoRegistro["res"], nil
	} else {
		return nil, err1
	}
}

func ConversorDoc(data string) (status interface{}, outputError interface{}) {
	return nil, nil
}

func ConvertHTMLToPDF(htmlData string) (io.Reader, error) {
	fmt.Println("Hola")
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		fmt.Println("Error al crear el generador de pdf: ", err)
		log.Fatal(err)
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(htmlData))))

	err = pdfg.Create()
	if err != nil {
		fmt.Println("Error al crear el pdf: ", err)
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./doc.pdf")
	if err != nil {
		fmt.Println("Error al escribir el pdf: ", err)
		log.Fatal(err)
	}

	// Convert buffer to io.Reader
	pdf := bytes.NewReader(pdfg.Bytes())

	return pdf, nil
}
