package controllers

import (
	"encoding/json"

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
	c.Mapping("ConstruirPlantilla", c.ConstruirPlantilla) // Nuevo método mapeado
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

	var body map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, err)
		c.ServeJSON()
		return
	}

	uid, err := services.ConstruirPlantilla(body)
	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, map[string]string{"uid": uid})
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, err)
	}

	c.ServeJSON()
}
