# Proyecto Synchronizer

Este repositorio contiene un proyecto en Go que utiliza la biblioteca GoFlow para crear y ejecutar flujos de trabajo.

## Descripción

Este proyecto se trata de una prueba de concepto para hacer pruebas con la librería GoFlow. Está diseñado para ejecutar un flujo de trabajo simple que realiza una tarea programada utilizando la biblioteca GoFlow.

El flujo de trabajo está definido en el archivo `main.go` y consiste en una tarea que imprime "hello" cada minuto.

## Instalación

1. Asegúrate de tener Go instalado en tu sistema.
2. Clona este repositorio en tu máquina local.
3. Ejecuta el siguiente comando para instalar las dependencias:

   ```bash
   go get github.com/fieldryand/goflow/v2
   ```

4. Ejecuta el siguiente comando
   ```bash
   go run main.go
   ```
5. Ve a `localhost:8181/ui` para ver el dashboard
