# plantillas_mid

Api intermediaria entre el cliente de Plantillas y las apis necesarios para la gestión del formato de los documentos de todo tipo.
El api principalmente es pensado para darle informacion a [plantillas_cliente](https://github.com/udistrital/plantillas_cliente). (esto no limita a su consumo desde el clinete o api que desee consumirle)  


## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)


### Variables de Entorno

# Ejemplo que se debe actualizar acorde al proyecto

PLANTILLAS_MID_RENDERIZADO_PLANTILLAS=[url del servicio de renderizado_plantillas_html]
PLANTILLAS_CRUD_SERVICE=[url del servicio de plantillas_crud_serverless]

**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con PLANTILLAS_MID...

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/plantillas_mid

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/plantillas_mid

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.

# 5. ejecutar
bee run
```

### Ejecución Dockerfile
```shell
# docker build --tag=plantillas_mid . --no-cache
# docker run -p 80:80 plantillas_mid
```


### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/plantillas_mid

#2. Moverse a la carpeta del repositorio
cd plantillas_mid

#3. Crear un fichero con el nombre **custom.env**
# En windows ejecutar:* ` ni custom.env`
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas

Pruebas unitarias
```shell
# Not Data
```

## Estado CI

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/platillas_mid/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/platillas_mid/) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/platillas_mid/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/platillas_mid/) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/platillas_mid/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/platillas_mid/) |



## Licencia

This file is part of platillas_mid.

platillas_mid is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

platillas_mid is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with novedades_crud. If not, see https://www.gnu.org/licenses/.
