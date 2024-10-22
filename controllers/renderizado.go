package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/udistrital/plantillas_mid/models"
	"github.com/udistrital/plantillas_mid/services"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/requestresponse"
)

type RenderizadoController struct {
	beego.Controller
}

func (c *RenderizadoController) RenderizarPlantilla() {
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
