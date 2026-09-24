package main

import (
	contadorvocales "Paquetes/ContadorVocales"
	"Paquetes/ConversorMonedas"
	"fmt"
)

func main() {
	fmt.Println("Que moneda desea cambiar: ")
	ConversorMonedas.Conversor()

	fmt.Println("Contador de Vocales")
	contadorvocales.Contador()
}
