package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/plantillas_mid/models"
	"github.com/udistrital/plantillas_mid/services"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/requestresponse"
)

type PlantillaController struct {
	beego.Controller
}

func (c *PlantillaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)
}

func (c *PlantillaController) Post() {
	defer errorhandler.HandlePanic(&c.Controller)

	var registroPlantilla models.Plantilla

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroPlantilla); err == nil {
		resultado, err := services.RegistrarPlantilla(registroPlantilla)

		if err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = requestresponse.APIResponseDTO(true, 201, resultado)
		} else {
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, err.Error())
		}
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, err.Error())
	}

	c.ServeJSON()
}

func (c *PlantillaController) Put() {
	defer errorhandler.HandlePanic(&c.Controller)

	idStr := c.Ctx.Input.Param(":id")
	var registroPlantilla models.Plantilla

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroPlantilla); err == nil {
		resultado, err := services.ActualizarPlantilla(idStr, registroPlantilla)

		if err == nil {
			c.Ctx.Output.SetStatus(200)
			c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
		} else {
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, err)
		}
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, err)
	}

	c.ServeJSON()
}

func (c *PlantillaController) RenderizarPlantilla() {
	defer errorhandler.HandlePanic(&c.Controller)

	var requestData models.Renderizado_plantilla
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestData)
	if err != nil {
		//log.Printf("Error al deserializar el JSON: %v", err)
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, "Error en el formato del JSON enviado")
		c.ServeJSON()
		return
	}

	uid, err := services.DefinirPlantilla(requestData)
	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, uid)
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, err.Error())
	}

	c.ServeJSON()
}
