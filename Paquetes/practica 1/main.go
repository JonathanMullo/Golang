package main

import (
	"fmt"
	"practica/operaciones"
	"practica/saludo"
)

func main() {
	fmt.Println("****Bienvenidos a la Clase de Paquetes****")
	mensaje := saludo.Saludar("Juan")
	fmt.Println(mensaje)

	fmt.Print("****Las Operaciones son las siguientes**** ")
	operaciones.Menu()
}
