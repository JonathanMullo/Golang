package main

import "fmt"

func main() {
	notas := [6][4]float64{
		{70.5, 80.5, 90.5, 90.0},
		{60.5, 81.5, 90.5, 85.5},
		{80.5, 95.5, 85.5, 75.5},
		{89.5, 92.5, 93.5, 89.5},
		{82.5, 84.7, 56.9, 98.6},
		{75.5, 85.6, 98.6, 85.6},
	}

	var sumaTotalClase float64
	totalCalificaciones := 6 * 4

	fmt.Println("El promedio de cada estudiante es:")

	for n := 0; n < 6; n++ {
		materiasEstudiante := notas[n][:]

		sumaEstudiante := 0.0
		notaAlta := materiasEstudiante[0]
		notaBaja := materiasEstudiante[0]

		for _, nota := range materiasEstudiante {
			sumaEstudiante += nota
			sumaTotalClase += nota

			if nota > notaAlta {
				notaAlta = nota
			}
			if nota < notaBaja {
				notaBaja = nota
			}
		} 

		promedioEstudiante := sumaEstudiante / float64(len(materiasEstudiante))

		fmt.Printf("Estudiante %d\n", n+1) 
		fmt.Printf("  *Promedio: %.2f\n", promedioEstudiante)
		fmt.Printf("  *Nota más alta: %.2f\n", notaAlta)
		fmt.Printf("  *Nota más baja: %.2f\n", notaBaja)

	} 

	promedioGeneral := sumaTotalClase / float64(totalCalificaciones)
	fmt.Printf("Promedio general de la clase: %2f\n ", promedioGeneral)
}