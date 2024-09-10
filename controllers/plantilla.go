package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/plantillas_mid/services"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/requestresponse"
)

type PlantillaController struct {
	beego.Controller
}

func (c *PlantillaController) URLMapping() {
	c.Mapping("DuplicarPlantilla", c.DuplicarPlantilla)
	c.Mapping("ConstruirPlantilla", c.ConstruirPlantilla)
	c.Mapping("PruebaConexion", c.PruebaConexion)

}

func (c *PlantillaController) DuplicarPlantilla() {
	defer errorhandler.HandlePanic(&c.Controller)

	idStr := c.Ctx.Input.Param(":id")
	plantillaDuplicada, err := services.DuplicarPlantilla(idStr)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, plantillaDuplicada)
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, err)
	}

	c.ServeJSON()
}

func (c *PlantillaController) ConstruirPlantilla() {
	defer errorhandler.HandlePanic(&c.Controller)

	// Ejemplo de datos a enviar
	html := "<html><body><h1>Hola, {{.Nombre}}</h1></body></html>"
	css := "h1 { color: red; }"
	datos := map[string]interface{}{
		"Nombre": "Mundo",
	}

	pdf, err := services.RenderizarPDF(html, css, datos)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = requestresponse.APIResponseDTO(false, 500, nil, err)
	} else {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, pdf)
	}

	c.ServeJSON()
}

func (c *PlantillaController) PruebaConexion() {
	err := services.ComprobarConexionCRUD()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Message": fmt.Sprintf("Error al conectar con el CRUD: %v", err),
		}
	} else {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{
			"Success": true,
			"Message": "Conexión con el CRUD exitosa.",
		}
	}
	c.ServeJSON()
}
