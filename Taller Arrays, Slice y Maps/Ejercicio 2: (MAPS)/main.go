package main

import "fmt"

func actividadGanadora(votos map[string]int) string {
	var ganadora string
	var mayor int

	for actividad, cantidad := range votos {
		if cantidad > mayor {
			mayor = cantidad
			ganadora = actividad
		}
	}
	return ganadora
}

func main() {
	votos := map[string]int{
		"Deportes":    0,
		"Videojuegos": 0,
		"Cine":        0,
		"Musica":      0,
	}

	fmt.Println("Elige una actividad:")
	fmt.Println("1. Deportes")
	fmt.Println("2. Videojuegos")
	fmt.Println("3. Cine")
	fmt.Println("4. Musica")

	for v := 1; v <= 5; v++ {
		var opcion string
		fmt.Print("Voto ", v, ": ")
		fmt.Scan(&opcion)

		switch opcion {
		case "1":
			votos["Deportes"]++
		case "2":
			votos["Videojuegos"]++
		case "3":
			votos["Cine"]++
		case "4":
			votos["Musica"]++
		}
	}
		ganadora := actividadGanadora(votos)
	fmt.Println("La actividad ganadora es", ganadora, "con", votos[ganadora], "votos")
}