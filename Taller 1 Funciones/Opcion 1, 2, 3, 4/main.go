package main

import "fmt"

// Opcion 1
func averageGrade(sumaNotas float64, cantidad int) float64 {
	if cantidad == 0 {
		return 0
	}
	return sumaNotas / float64(cantidad)
}

func ProcesarNotas() {
	var cantidad int
	fmt.Print("Ingrese la cantidad de estudiantes en el curso: ")
	fmt.Scan(&cantidad)

	var sumaNotas float64
	var nota float64

	for i := 1; i <= cantidad; i++ {
		fmt.Printf("Ingrese la nota del estudiante %d (0 a 100): ", i)
		fmt.Scan(&nota)
		sumaNotas += nota
	}
	promedio := averageGrade(sumaNotas, cantidad)
	fmt.Printf("\nEl promedio del curso es: %.2f\n", promedio)

	if promedio >= 70 {
		fmt.Println("Estado: Curso Aprobado")
	} else {
		fmt.Println("Estado: Curso Reprobado")
	}

	switch {
	case promedio >= 90 && promedio <= 100:
		fmt.Println("Rendimiento: Excellent performance")
	case promedio >= 80 && promedio < 90:
		fmt.Println("Rendimiento: Good performance")
	case promedio >= 70 && promedio < 80:
		fmt.Println("Rendimiento: Satisfactory performance")
	case promedio < 70:
		fmt.Println("Rendimiento: Needs improvement")
	default:
		fmt.Println("Nota fuera de rango normal.")
	}
}

// Opcion 2
func SumaHastaN() {
	var n int
	var sumatoriaNum int

	fmt.Println("Ingrese un numero:")
	fmt.Scan(&n)

	for d := 1; d <= n; d++ {
		sumatoriaNum += d
	}
	fmt.Println("La suma del 1 al", n, "es:", sumatoriaNum)
}

// Opcion 3
func CelsiusFahrenheit() {
	var celsius float64

	fmt.Println("Ingrese la temperatura que desea convertir a Fahrenheit:")
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 1.8) + 32
	fmt.Printf("%.2f Celsius a %.2f Fahrenheit\n", celsius, fahrenheit)
}

// Opción 4
func FahrenheitCelsius() {
	var fahrenheit float64

	fmt.Println("Ingresa la temperatura que desea convertir a Celsius:")
	fmt.Scan(&fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f Fahrenheit a %.2f Celsius\n", fahrenheit, celsius)
}

//menu
func main() {
	var opcion string

	for {
		fmt.Println("\n Menu De Calculos ")
		fmt.Println("1. Promedio de Estudiantes")
		fmt.Println("2. Suma hasta N")
		fmt.Println("3. Celsius a Fahrenheit")
		fmt.Println("4. Fahrenheit a Celsius")
		fmt.Println("0. Salir (o escribe 'salir')")
		fmt.Print("Elija una opción: ")
		fmt.Scan(&opcion)
		if opcion == "0" || opcion == "salir" || opcion == "Salir" {
			break
		}

		switch opcion {
		case "1":
			ProcesarNotas()
		case "2":
			SumaHastaN()
		case "3":
			CelsiusFahrenheit()
		case "4":
			FahrenheitCelsius()
		default:
			fmt.Println("Opción no válida.")
		}
	}
}
