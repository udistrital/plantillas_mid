package controllers

import (
	"encoding/json"
	"log"

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
	c.Mapping("PruebaConexion", c.ConstruirDocumento)
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

	var requestData map[string]interface{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestData)
	if err != nil {
		log.Printf("Error al deserializar el JSON: %v", err)
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, "Error en el formato del JSON enviado")
		c.ServeJSON()
		return
	}

	uid, err := services.ProcesarPlantilla(requestData)
	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, uid)
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, err.Error())
	}

	c.ServeJSON()
}

func (c *PlantillaController) ConstruirDocumento() {
	defer errorhandler.HandlePanic(&c.Controller)

	var requestData map[string]interface{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestData)
	if err != nil {
		log.Printf("Error al deserializar el JSON: %v", err)
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, "Error en el formato del JSON enviado")
		c.ServeJSON()
		return
	}

	base64PDF, err := services.ProcesarDocumento(requestData)
	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, base64PDF)
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, err.Error())
	}

	c.ServeJSON()
}
