package contadorvocales

import (
	"bufio"
	"fmt"
	"os"
)

func Contador() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Ingrese una palabra o frase: ")

	scanner.Scan()
	frase := scanner.Text()

	a, e, i, o, u := 0, 0, 0, 0, 0

	for _, letra := range frase {
		switch letra {
		case 'a', 'A', 'á', 'Á':
			a++
		case 'e', 'E', 'é', 'É':
			e++
		case 'i', 'I', 'í', 'Í':
			i++
		case 'o', 'O', 'ó', 'Ó':
			o++
		case 'u', 'U', 'ú', 'Ú':
			u++
		}
	}

	fmt.Println("\n--- Conteo de Vocales ---")
	fmt.Println("A:", a)
	fmt.Println("E:", e)
	fmt.Println("I:", i)
	fmt.Println("O:", o)
	fmt.Println("U:", u)
}
