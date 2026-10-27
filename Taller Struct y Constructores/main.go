package main
import "fmt"

type Producto struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{id, nombre, precio, cantidad}
}
func (p *Producto) ActualizarPrecio(nuevoPrecio float64) {
	p.Precio = nuevoPrecio
}

func (p *Producto) ActualizarCantidad(nuevaCantidad int) {
	p.Cantidad = nuevaCantidad
}

func (p Producto) MostrarInfo() {
	fmt.Printf("ID: %d , %s , $%.2f , Cantidad: %d\n", p.ID, p.Nombre, p.Precio, p.Cantidad)
}

type Inventario struct {
	Productos map[int]Producto
}

func (inv *Inventario) AgregarProducto(p Producto) {
	inv.Productos[p.ID] = p
}

func (inv *Inventario) EliminarProducto(id int) {
	delete(inv.Productos, id)
}

func (inv *Inventario) BuscarProducto(id int) (Producto, bool) {
	p, existe := inv.Productos[id]
	return p, existe
}

func (inv *Inventario) ListarProductos() {
	for _, p := range inv.Productos {
		p.MostrarInfo()
	}
}

func main() {
	inv := Inventario{Productos: make(map[int]Producto)}

	p1 := NuevoProducto(1, "Laptop", 850.50, 10)
	p2 := NuevoProducto(2, "Mouse", 15.99, 50)
	p2.ActualizarPrecio(18.50)
	p2.ActualizarCantidad(45)

	inv.AgregarProducto(p1)
	inv.AgregarProducto(p2)

	fmt.Println("Lista de productos:")
	inv.ListarProductos()

	if p, existe := inv.BuscarProducto(1); existe {
		fmt.Println("Encontrado:")
		p.MostrarInfo()
	}

	inv.EliminarProducto(1)
	fmt.Println("Después de eliminar:")
	inv.ListarProductos()
}