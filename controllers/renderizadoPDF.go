package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/plantillas_mid/models"
	"github.com/udistrital/plantillas_mid/services"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/requestresponse"
)

type RenderizadoPDFController struct {
	beego.Controller
}

// URLMapping ...
func (c *RenderizadoPDFController) URLMapping() {
	c.Mapping("RenderizarPDF", c.RenderizarPDF)
}

// RenderizarPDF ...
// @Title RenderizarPDF
// @Description Genera un PDF (base64) a partir de contenido HTML y datos personalizados para su reemplazo
// @Param	body		body 	models.RenderizadoPDF	true		"body for RenderizadoPDF content"
// @Success 201 {int} models.Rol
// @Failure 400 body is empty
// @router /renderizar-pdf [post]
func (c *RenderizadoPDFController) RenderizarPDF() {
	defer errorhandler.HandlePanic(&c.Controller)

	var requestData models.RenderizadoPDF
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestData)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, "Error en el formato del JSON enviado")
		c.ServeJSON()
		return
	}

	plantillaRenderizada, err := services.RenderizarPDF(requestData)
	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, plantillaRenderizada)
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = requestresponse.APIResponseDTO(false, 400, nil, err.Error())
	}

	c.ServeJSON()
}
