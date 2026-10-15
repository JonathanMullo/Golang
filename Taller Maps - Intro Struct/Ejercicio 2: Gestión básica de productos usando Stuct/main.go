package main
import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func calcularValorTotal(p Producto) float64 {
	return p.Precio * float64(p.Cantidad)
}
func tieneStock(p Producto) bool {
	return p.Cantidad > 0
}

func mostrarProducto(p Producto) {
	estado := "Disponible"
	if !tieneStock(p) {
		estado = "Sin stock"
	}
	total := calcularValorTotal(p)
	fmt.Printf("Producto: %-6s, Precio: $%5.2f, Stock:%3d, Total: $%6.2f, Estado:%s\n",
		p.Nombre, p.Precio, p.Cantidad, total, estado)
}

func main() {
	p1 := Producto{Nombre: "Teclado", Precio: 45.50, Cantidad: 10}
	p2 := Producto{Nombre: "Mouse", Precio: 15.00, Cantidad: 0}
	p3 := Producto{Nombre: "Monitor", Precio: 160.00, Cantidad: 4}
	listaProductos := []Producto{p1, p2, p3}

	fmt.Println(" GESTION DE PRODUCTOS ")
	for _, prod := range listaProductos {
		mostrarProducto(prod)
	}
}