package controllers

import (
	"encoding/json"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/udistrital/plantillas_mid/models"
)

// MinutaController operations for Minuta
type MinutaController struct {
	beego.Controller
}

// URLMapping ...
func (c *MinutaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Minuta
// @Param	body		body 	models.Minuta	true		"body for Minuta content"
// @Success 201 {object} models.Minuta
// @Failure 403 body is empty
// @router / [post]
func (c *MinutaController) Post() {
	var registroMinuta map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroMinuta); err == nil {

		result, err1 := models.RegistrarMinuta(registroMinuta)

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
// @Description get Minuta by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Minuta
// @Failure 403 :id is empty
// @router /:id [get]
func (c *MinutaController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	var minuta map[string]interface{}

	var alerta models.Alert

	minuta, err := models.ConsultarMinuta(idStr)

	fmt.Println(err)
	if err == nil {
		alerta.Type = "OK"
		alerta.Code = "200"
		alerta.Body = minuta
	} else {
		alerta.Type = "error"
		alerta.Code = "400"
		alerta.Body = "No se ha podido realizar la petición GET"
	}

	c.Data["json"] = alerta
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Minuta
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Minuta
// @Failure 403
// @router / [get]
func (c *MinutaController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Minuta
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Minuta	true		"body for Minuta content"
// @Success 200 {object} models.Minuta
// @Failure 403 :id is not int
// @router /:id [put]
func (c *MinutaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	var registroMinuta map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroMinuta); err == nil {

		result, err1 := models.ActualizarSeccion(idStr, registroMinuta)

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
// @Description delete the Minuta
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *MinutaController) Delete() {

}
