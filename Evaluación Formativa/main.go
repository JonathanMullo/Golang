package main

import (
	"fmt"
	"time"
)

type Medicine struct {
	name            string
	manufacturer    string
	manufactureDate time.Time
	shelfLife       int
}

func NuevaMedicine(name string, manufacturer string, manufactureDate time.Time, shelfLife int) Medicine {
	m := Medicine{}
	m.SetName(name)
	m.SetManufacturer(manufacturer)
	m.SetManufactureDate(manufactureDate)
	m.SetShelfLife(shelfLife)
	return m
}

func (m Medicine) GetName() string {
	return m.name
}

func (m Medicine) GetManufacturer() string {
	return m.manufacturer
}

func (m Medicine) GetManufactureDate() time.Time {
	return m.manufactureDate
}

func (m Medicine) GetShelfLife() int {
	return m.shelfLife
}

func (m *Medicine) SetName(nuevoName string) {
	if nuevoName != "" {
		m.name = nuevoName
	} else {
		fmt.Println("Error: El nombre no puede estar vacío.")
	}
}

func (m *Medicine) SetManufacturer(nuevaManufacturer string) {
	if nuevaManufacturer != "" {
		m.manufacturer = nuevaManufacturer
	} else {
		fmt.Println("Error: El fabricante no debe estar vacío.")
	}
}

func (m *Medicine) SetManufactureDate(nuevaFecha time.Time) {
	if !nuevaFecha.IsZero() {
		m.manufactureDate = nuevaFecha
	} else {
		fmt.Println("Error: La fecha de fabricación no es válida.")
	}
}

func (m *Medicine) SetShelfLife(nuevaShelfLife int) {
	if nuevaShelfLife > 0 {
		m.shelfLife = nuevaShelfLife
	} else {
		fmt.Println("Error: La vida útil debe ser mayor a 0.")
	}
}

func (m Medicine) CalcularCaducidad() time.Time {
	return m.manufactureDate.AddDate(0, m.shelfLife, 0)
}

type Tablet struct {
	Medicine
	dosePerTablet  float64
	isPrescription bool
}

func NuevaTablet(name string, manufacturer string, manufactureDate time.Time, shelfLife int, dosePerTablet float64, isPrescription bool) Tablet {
	t := Tablet{
		Medicine:       NuevaMedicine(name, manufacturer, manufactureDate, shelfLife),
		isPrescription: isPrescription,
	}
	t.SetDosePerTablet(dosePerTablet)
	return t
}

func (t Tablet) GetDosePerTablet() float64 {
	return t.dosePerTablet
}

func (t Tablet) GetIsPrescription() bool {
	return t.isPrescription
}

func (t *Tablet) SetDosePerTablet(dosePerTablet float64) {
	if dosePerTablet > 0 {
		t.dosePerTablet = dosePerTablet
	} else {
		fmt.Println("Error: La dosis por tableta debe ser mayor a 0.")
	}
}

func (t *Tablet) SetIsPrescription(isPrescription bool) {
	t.isPrescription = isPrescription
}

func (t Tablet) MostrarDetalles() {
	receta := "No"
	if t.isPrescription {
		receta = "Sí"
	}
	fmt.Println("--- DETALLE TABLETA ---")
	fmt.Printf("Nombre: %s\n", t.GetName())
	fmt.Printf("Fabricante: %s\n", t.GetManufacturer())
	fmt.Printf("Fecha Fabricación: %s\n", t.GetManufactureDate().Format("2006-01-02"))
	fmt.Printf("Vida Útil: %d meses\n", t.GetShelfLife())
	fmt.Printf("Fecha Caducidad: %s\n", t.CalcularCaducidad().Format("2006-01-02"))
	fmt.Printf("Dosis por Tableta: %.2f mg\n", t.GetDosePerTablet())
	fmt.Printf("Requiere Receta: %s\n\n", receta)
}

type Syrup struct {
	Medicine
	volume float64
	flavor string
}

func NuevoSyrup(name string, manufacturer string, manufactureDate time.Time, shelfLife int, volume float64, flavor string) Syrup {
	s := Syrup{
		Medicine: NuevaMedicine(name, manufacturer, manufactureDate, shelfLife),
		flavor:   flavor,
	}
	s.SetVolume(volume)
	return s
}

func (s Syrup) GetVolume() float64 {
	return s.volume
}

func (s Syrup) GetFlavor() string {
	return s.flavor
}

func (s *Syrup) SetVolume(volume float64) {
	if volume > 0 {
		s.volume = volume
	} else {
		fmt.Println("Error: El volumen debe ser mayor a 0.")
	}
}

func (s *Syrup) SetFlavor(flavor string) {
	s.flavor = flavor
}

func (s Syrup) MostrarDetalles() {
	fmt.Println("Detalle del Jarabe")
	fmt.Printf("Nombre: %s\n", s.GetName())
	fmt.Printf("Fabricante: %s\n", s.GetManufacturer())
	fmt.Printf("Fecha Fabricación: %s\n", s.GetManufactureDate().Format("2006-01-02"))
	fmt.Printf("Vida Útil: %d meses\n", s.GetShelfLife())
	fmt.Printf("Fecha Caducidad: %s\n", s.CalcularCaducidad().Format("2006-01-02"))
	fmt.Printf("Volumen: %.2f ml\n", s.GetVolume())
	fmt.Printf("Sabor: %s\n\n", s.GetFlavor())
}

func main() {
	fecha1 := time.Date(2023, time.January, 15, 0, 0, 0, 0, time.UTC)
	fecha2 := time.Date(2024, time.March, 10, 0, 0, 0, 0, time.UTC)

	inventarioTabletas := []Tablet{
		NuevaTablet("Paracetamol", "Genfar", fecha1, 24, 500.0, false),
	}
	inventarioJarabes := []Syrup{
		NuevoSyrup("Ambroxol", "MK", fecha2, 12, 120.0, "Frambuesa"),
	}

	fmt.Println("Inventario")
	for _, tab := range inventarioTabletas {
		tab.MostrarDetalles()
	}
	for _, syr := range inventarioJarabes {
		syr.MostrarDetalles()
	}

	// 2. Modificar con setter y verificar con getter
	fmt.Println("Actualizacion")
	fmt.Printf("Dosis previa: %.2f mg\n", inventarioTabletas[0].GetDosePerTablet())
	inventarioTabletas[0].SetDosePerTablet(650.0)
	fmt.Printf("Dosis actualizada: %.2f mg\n\n", inventarioTabletas[0].GetDosePerTablet())

	// 3. Mostrar cálculo de fecha de caducidad
	fmt.Println("Fecha de Caduciadad")
	fmt.Printf("Tableta '%s' caduca el: %s\n",
		inventarioTabletas[0].GetName(),
		inventarioTabletas[0].CalcularCaducidad().Format("2006-01-02"))

	fmt.Printf("Jarabe '%s' caduca el: %s\n",
		inventarioJarabes[0].GetName(),
		inventarioJarabes[0].CalcularCaducidad().Format("2006-01-02"))
}





		