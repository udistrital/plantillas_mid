package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plantillas_mid/models"
	"github.com/udistrital/plantillas_mid/services"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/requestresponse"
)

type RenderizadoController struct {
	beego.Controller
}

// URLMapping ...
func (c *RenderizadoController) URLMapping() {
	c.Mapping("Renderizar", c.Renderizar)
}

// Renderizar ...
// @Title Renderizar
// @Description Genera un PDF (base64) utilizando un ID de plantilla y datos personalizados para su reemplazo
// @Param	body		body 	models.Renderizado	true		"body for Renderizado content"
// @Success 201 {int} models.Rol
// @Failure 400 body is empty
// @router /renderizar [post]
func (c *RenderizadoController) Renderizar() {
	defer errorhandler.HandlePanic(&c.Controller)
	logs.Info("-------Entra controller renderizado")
	var requestData models.Renderizado
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestData)
	if err != nil {
		logs.Info("ERROOOOOOR controller handlepanic")
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, "Error en el formato del JSON enviado")
		c.ServeJSON()
		return
	}

	plantillaRenderizada, err := services.Renderizar(requestData)
	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, plantillaRenderizada)
	} else {
		logs.Info("ERROOOOOOR controller renderizado")
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, err.Error())
	}

	c.ServeJSON()
}
