package ConversorMonedas

import "fmt"

func euros(dolar float64) float64 {
	euro := 0.88
	return dolar * euro
}
func libras(dolar float64) float64 {
	libra := 0.7557
	return dolar * libra

}
func won(dolar float64) float64 {
	won := 1.365
	return dolar + won

}
func btc(dolar float64) float64 {
	btc := 0.000012
	return dolar + btc
}

func Conversor() {
	var opcion string

	for {
		fmt.Println("\n Menu De Conversor ")
		fmt.Println("1. Dolar a Euro")
		fmt.Println("2. Dolar a Libra")
		fmt.Println("3. Dolar a Won")
		fmt.Println("4. Dolar a BTC")
		fmt.Println("0. Salir (o escribe 'salir')")
		fmt.Print("Elija una opción: ")
		fmt.Scan(&opcion)
		if opcion == "0" || opcion == "salir" || opcion == "Salir" {
			break
		}
		if opcion == "1" || opcion == "2" || opcion == "3" || opcion == "4" {
			var dolares float64
			fmt.Println("Ingrese la cantidad de Dolares que desea conversir (USD): ")
			fmt.Scan(&dolares)

			switch opcion {
			case "1":
				resultado := euros(dolares)
				fmt.Println(resultado)
			case "2":
				resultado := libras(dolares)
				fmt.Println(resultado)
			case "3":
				resultado := won(dolares)
				fmt.Println(resultado)
			case "4":
				resultado := btc(dolares)
				fmt.Println(resultado)
			default:
				fmt.Println("Opción no válida.")
			}
		}
	}
}
