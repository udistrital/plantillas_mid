package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/plantillas_mid/models"
)

// PlantillaController operations for Plantilla
type PlantillaController struct {
	beego.Controller
}

// URLMapping ...
func (c *PlantillaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Plantilla
// @Param	body		body 	models.Plantilla	true		"body for Plantilla content"
// @Success 201 {object} models.Plantilla
// @Failure 403 body is empty
// @router / [post]
func (c *PlantillaController) Post() {
	var registroPlantilla map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroPlantilla); err == nil {

		result, err1 := models.RegistrarPlantilla(registroPlantilla)

		if err1 == nil {
			alertErr.Type = "OK"
			alertErr.Code = "200"
			alertErr.Body = result
		} else {
			alertErr.Type = "error"
			alertErr.Code = "400"
			alertas = append(alertas, err1)
			alertErr.Body = alertas
			c.Ctx.Output.SetStatus(400)
		}

	} else {
		alertErr.Type = "error"
		alertErr.Code = "400"
		alertas = append(alertas, err.Error())
		alertErr.Body = alertas
		c.Ctx.Output.SetStatus(400)
	}

	c.Data["json"] = alertErr
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Plantilla by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Plantilla
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PlantillaController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Plantilla
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Plantilla
// @Failure 403
// @router / [get]
func (c *PlantillaController) GetAll() {
	var resultado interface{}
	//var Result interface{}
	var alerta models.Alert
	//resultado = models.GetPlantillas()
	//fmt.Println(errResultado, resultado)
	alerta.Type = "OK"
	alerta.Code = "200"
	alerta.Body = resultado
	c.Data["json"] = alerta
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Plantilla
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Plantilla	true		"body for Plantilla content"
// @Success 200 {object} models.Plantilla
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PlantillaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	var registroNovedad map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroNovedad); err == nil {

		result, err1 := models.ActualizarPlantilla(idStr)

		if err1 == nil {
			alertErr.Type = "OK"
			alertErr.Code = "200"
			alertErr.Body = result
		} else {
			alertErr.Type = "error"
			alertErr.Code = "400"
			alertas = append(alertas, err1)
			alertErr.Body = alertas
			c.Ctx.Output.SetStatus(400)
		}

	} else {
		alertErr.Type = "error"
		alertErr.Code = "400"
		alertas = append(alertas, err.Error())
		alertErr.Body = alertas
		c.Ctx.Output.SetStatus(400)
	}

	c.Data["json"] = alertErr
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Plantilla
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PlantillaController) Delete() {

}
