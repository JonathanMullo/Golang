package main

import "fmt"
// Opcion 1
func 

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
    var opcion int

    fmt.Println(" Menu De Calculos")
    fmt.Println("2. Suma hasta N")
    fmt.Println("3. Celsius a Fahrenheit")
    fmt.Println("4. Fahrenheit a Celsius")
    fmt.Print("Elija una opción: ")
    fmt.Scan(&opcion)

    switch opcion {
    case 2:
        SumaHastaN()
    case 3:
        CelsiusFahrenheit()
    case 4:
        FahrenheitCelsius()
    default: 
        fmt.Println("Opción no válida.")
    }
}