package main
import "fmt"

func main() {
	conteoPalabra := make(map[string]int)
	fmt.Println("Ingrese una frase o texto:")
	fmt.Println("Escriba 'fin' para terminar el proceso:")

	var palabra string
	for {
		fmt.Scan(&palabra)

		if palabra == "fin" {
			break
		}
		conteoPalabra[palabra]++
	}

	fmt.Println("\nConteo de Palabras")
	for p, cantidad := range conteoPalabra {
		fmt.Println(p, ":", cantidad)
	}

	palabraMayor := ""
	maximo := 0

	for p, cantidad := range conteoPalabra {
		if cantidad > maximo {
			maximo = cantidad
			palabraMayor = p
		}
	}

	if maximo > 0 {
		fmt.Printf("\nLa palabra que más se repite es '%s' con %d veces.\n", palabraMayor, maximo)
	} else {
		fmt.Println("\nNo se ingresaron palabras.")
	}
}