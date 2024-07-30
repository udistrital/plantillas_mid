package models

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html"
	"html/template"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/jung-kurt/gofpdf"

	// "github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
)

type Plantilla struct {
	Tipo              string
	Nombre            string
	Descripcion       string
	Contenido         string
	Secciones         []Seccion
	EnlaceDoc         string
	FechaCreacion     time.Time
	FechaModificacion time.Time
	version           float64
	versionActual     bool
}

type PlantillaPost struct {
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

func RegistrarPlantilla(plantilla Plantilla) (string, interface{}) {

	fmt.Println("Función RegistrarPlantilla")

	var PlantillaPost PlantillaPost
	// var resultadoRegistro map[string]interface{}
	// var errRegPlantilla interface{}

	// dataPlantilla := plantilla.Nombre

	// result1, errConvert := ConversorDoc(dataPlantilla)
	// fmt.Println("Result1: ", result1+errConvert.Error())
	result, errConvert := ConversorHTML(plantilla)
	// errConvert := generatePDF(dataPlantilla, "C:/Users/Youssef/Desktop/test.pdf")
	if errConvert != nil {
		fmt.Println("Error al convertir la estructura a HTML ", errConvert)
		return "nil", errConvert
	} else {
		fmt.Println("Pdf generado correctamente!")

		fmt.Println("Result: ", result)

		PlantillaPost.Tipo = plantilla.Tipo
		PlantillaPost.Nombre = plantilla.Nombre
		PlantillaPost.Descripcion = plantilla.Descripcion
		PlantillaPost.EnlaceDoc = plantilla.EnlaceDoc
		PlantillaPost.FechaCreacion = plantilla.FechaCreacion
		PlantillaPost.FechaModificacion = plantilla.FechaModificacion
		PlantillaPost.version = plantilla.version
		PlantillaPost.versionActual = plantilla.versionActual

		PlantillaPost.Contenido = result

		fmt.Println("PlantillaPost: ", PlantillaPost)

		// errRegPlantilla = request.SendJson(beego.AppConfig.String("PlantillasCrudService")+"/plantilla", "POST", &resultadoRegistro, PlantillaPost)

		// if resultadoRegistro["Status"] == "400" || errRegPlantilla != nil {
		// 	fmt.Println(errRegPlantilla)
		// 	return "nil", errRegPlantilla

		// } else {
		// 	formatdata.JsonPrint(resultadoRegistro)
		// 	return "nil", nil

		// }
		return result, nil
	}
	return "nil", nil
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

func ConversorDoc(data string) (string, error) {

	pdf := gofpdf.New("P", "mm", "A4", "")

	tmpl := template.Must(template.New("").Parse(data))

	var buf bytes.Buffer
	err := tmpl.Execute(&buf, nil)
	if err != nil {
		return "nil", err
	}
	pdf.SetFont("Arial", "", 12)
	pdf.AddPage()

	for _, line := range strings.Split(buf.String(), "\n") {
		// Eliminar espacios adicionales
		line = strings.TrimSpace(line)

		// Verificar si la línea es un párrafo HTML o una línea nueva
		if strings.HasPrefix(line, "<p>") && strings.HasSuffix(line, "</p>") {
			// Extraer el contenido del párrafo
			content := strings.TrimPrefix(line, "<p>")
			content = strings.TrimSuffix(content, "</p>")

			// Agregar el párrafo al PDF
			pdf.MultiCell(0, 10, content, "", "C", false)
		} else if line == "<br>" {
			// Agregar una nueva línea al PDF para <br>
			pdf.Ln(10)
		}
	}

	err = pdf.Output(&buf)
	if err != nil {
		return "nil", err
	}

	err = pdf.OutputFileAndClose("C:/Users/Youssef/Desktop/test.pdf")
	if err != nil {
		return "nil", err
	}

	pdfBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())

	return pdfBase64, nil
}

func generatePDF(htmlString, outputPath string) error {
	// Crear un nuevo documento PDF
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Convertir el HTML a PDF usando una plantilla
	// tmpl, err := template.New("html").Parse(htmlString)
	// if err != nil {
	// 	return err
	// }

	// // Buffer para almacenar el HTML renderizado
	// var htmlBuffer bytes.Buffer

	// // Renderizar el HTML en el buffer
	// if err := tmpl.Execute(&htmlBuffer, nil); err != nil {
	// 	return err
	// }
	text := html.UnescapeString(htmlString)
	fmt.Println("Text: ", text)

	pdf.SetFont("Arial", "", 12)

	// Agregar una nueva página al PDF
	pdf.AddPage()

	// Configurar el tamaño y la posición de la celda
	pdf.SetFontSize(12)
	pdf.CellFormat(190, 10, "", "", 1, "C", false, 0, "")

	// fmt.Println("HtmlBuffer: ", htmlBuffer.String())
	// fmt.Println("Length: ", len(htmlBuffer.String()))
	fmt.Println("HtmlBuffer: ", text)
	fmt.Println("Length: ", len(text))

	pdf.SetFontSize(12)
	pdf.SetFont("Arial", "", 12)

	// Escribir el HTML renderizado en el PDF
	pdf.MultiCell(190, 10, text, "", "C", false)

	// Guardar el PDF en el archivo de salida
	return pdf.OutputFileAndClose(outputPath)
}

func ConversorHTML(plantilla Plantilla) (string, error) {

	fmt.Println("Entra a la función ConversorHTML")

	fmt.Println("Plantilla: ", plantilla)

	texto := plantilla.Contenido

	fmt.Println("Texto: ", texto)

	patron := `\[_\]`

	expReg := regexp.MustCompile(patron)

	resultados := expReg.Split(texto, -1)

	arreglo := make([]string, 0)

	for i, v := range resultados {
		arreglo = append(arreglo, v)
		if i < len(resultados)-1 {
			arreglo = append(arreglo, "[_]")
		}
	}

	fmt.Println(arreglo)

	var seccciones = plantilla.Secciones

	data := map[string]interface{}{
		"Title": plantilla.Nombre,
		"Items": []Seccion{seccciones[0], seccciones[1], seccciones[2], seccciones[3]},
	}

	// Crear una nueva plantilla desde el string
	tmpl, err := template.New("miPlantilla").Parse(plantilla.Nombre)
	if err != nil {
		return "", err
	}

	// Ejecutar la plantilla con los datos
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		return "", err
	}
	return "nil", nil
}
