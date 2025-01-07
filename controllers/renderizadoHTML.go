package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/plantillas_mid/models"
	"github.com/udistrital/plantillas_mid/services"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/requestresponse"
)

type RenderizadoHTMLController struct {
	beego.Controller
}

// URLMapping ...
func (c *RenderizadoHTMLController) URLMapping() {
	c.Mapping("RenderizarHTML", c.RenderizarHTML)
}

// RenderizarHTML ...
// @Title RenderizarHTML
// @Description Genera un HTML utilizando un ID de plantilla y datos personalizados para su reemplazo
// @Param	body		body 	models.RenderizadoHTML	true		"body for RenderizadoHTML content"
// @Success 201 {int} models.Rol
// @Failure 400 body is empty
// @router /renderizar-html [post]
func (c *RenderizadoHTMLController) RenderizarHTML() {
	defer errorhandler.HandlePanic(&c.Controller)

	var requestData models.RenderizadoHTML
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestData)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, "Error en el formato del JSON enviado")
		c.ServeJSON()
		return
	}

	plantillaRenderizada, err := services.RenderizarHTML(requestData)
	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, plantillaRenderizada)
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, err.Error())
	}

	c.ServeJSON()
}
