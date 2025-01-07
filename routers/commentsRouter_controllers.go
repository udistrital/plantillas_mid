package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:RenderizadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:RenderizadoController"],
        beego.ControllerComments{
            Method: "Renderizar",
            Router: "/renderizar",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:RenderizadoHTMLController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:RenderizadoHTMLController"],
        beego.ControllerComments{
            Method: "RenderizarHTML",
            Router: "/renderizar-html",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:RenderizadoPDFController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:RenderizadoPDFController"],
        beego.ControllerComments{
            Method: "RenderizarPDF",
            Router: "/renderizar-pdf",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
