package main

import "fmt"
/*

func <nombre>(param1, param2, ....param n)<valores de retorno>{
   ----------------------
   ----------------------
   ----------------------
   //return en el caso de que nuestra funcion retorne valores 
} 

*/

func saludar(){
	fmt.Println("Hola esta es mi primera funcion")

}

func bienvenida(nombre string){
	fmt.Println("Bienvenid@", nombre)
}
	
func suma(a, b int) int {
    return a + b
}

func resta(a, b int) int {
    return a - b
}
func SumaResta(s1, s2 int) (int, int) {
	var ResSuma int = s1 + s2
	var ResResta int

	if s2 > s1 {
		ResResta = 0
	} else {
		ResResta = s1 - s2
	}

	return ResSuma, ResResta
}

/*Variadic Functiones ======= Funciones Variadicas*/

func mostrarNum(numeros ...int){
	fmt.Println("Los numeros ingresados son: ", numeros)
}
func SumMostrarNum (numeros ...int) int {
	total :=0
	for _,numeros:= range numeros{
		total += numeros
	}
	return total
}
func main() {
	var usr string

	fmt.Println("Ingrese su nombre:")
	fmt.Scan(&usr)
	saludar()
	bienvenida(usr)

	fmt.Println("El resultado de la suma es:", suma(4, 5))
	fmt.Println("El resultado de la resta es:", resta(8, 3))

	ResSuma, ResResta := SumaResta(3, 4)

	fmt.Println("Suma:", ResSuma)
	fmt.Println("Resta:", ResResta)

	mostrarNum(5, 10, 15, 20)

	fmt.Println("La sumatoria es: ", SumMostrarNum (1, 2 ,3 ,4 ,5 ,6 ,7 ,8 ,9))
}



