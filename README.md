
# Chat Go (WebSocket)

## Descripción del Proyecto

El objetivo de este proyecto es desarrollar un servidor en el lenguaje de programación Go que tenga funcionamiento de chat en tiempo real usando WebSockets, donde los mensajes se almacenan en una base de datos MongoDB.

### Requisitos Previos

- Instalación de Go: [Descargar e instalar Go](https://golang.org/doc/install)

### MongoDB usando MongoDB Atlas
Desde la pagina oficial.

### Configuración del Proyecto

1. Clone este repositorio en su máquina local:

   ```bash
   git clone https://github.com/jerapepe/Server-WS-MG.git
   ```

2. Navegue al directorio del proyecto:

   ```bash
   cd Server-WS-MG
   ```

3. Instale las dependencias del proyecto:

   ```bash
   go mod tidy
   ```


## Uso

Para ejecutar el servidor, simplemente ejecute el siguiente comando desde el directorio cmd del proyecto:

```bash
go run main.go
```

El servidor comenzará a escuchar en el puerto especificado.

## Licencia

Este proyecto está bajo la licencia [MIT License](LICENSE).

---
