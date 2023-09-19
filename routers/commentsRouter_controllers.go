package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:EstiloFuenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:EstiloFuenteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:EstiloFuenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:EstiloFuenteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:EstiloFuenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:EstiloFuenteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:EstiloFuenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:EstiloFuenteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:EstiloFuenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:EstiloFuenteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:MinutaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:MinutaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:MinutaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:MinutaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:MinutaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:MinutaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:MinutaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:MinutaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:MinutaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:MinutaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:PlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:PlantillaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:PlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:PlantillaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:PlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:PlantillaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:PlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:PlantillaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:PlantillaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:PlantillaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:SeccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:SeccionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:SeccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:SeccionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:SeccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:SeccionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:SeccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:SeccionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:SeccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:SeccionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:TituloController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:TituloController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:TituloController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:TituloController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:TituloController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:TituloController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:TituloController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:TituloController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:TituloController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plantillas_mid/controllers:TituloController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
