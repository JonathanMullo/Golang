package main

import ("fmt"
		"time")

type Medicine struct {
	name string
	manufacturer string
	manufactureDate time.Time
	shelfLife int
}

func NuevaMedicine(name string, manufacturer string, manufactureDate time.Time, shelfLife int) Medicine {
	return Medicine{
		name: name,
		manufacturer: manufacturer,
		manufactureDate: manufactureDate,
		shelfLife: shelfLife,
	}
}
func (m Medicine) GetName() string {
	return m.name
}
func (m Medicine) GetManufacturer() string{
	return m.manufacturer
}
func (m Medicine) GetManufacturerDate() time.Time{
	return m.manufactureDate
}
func (m Medicine) GetShelfLife() int {
	return m.shelfLife
}

func (m *Medicine) SetName(nuevoName string) {
	if nuevoName != "" {
		m.name = nuevoName
	} else {
		fmt.Println("El nombre no puede estar vacío")
	}
}
func (m *Medicine) SetManufacturer(nuevaManufacturer string) {
	if nuevaManufacturer != "" {
		m.manufacturer = nuevaManufacturer
	} else {
		fmt.Println("El fabricante no debe estar vacío")
	}
}

func (m *Medicine) SetManufactureDate(nuevaFecha time.Time) {
	if !nuevaFecha.IsZero() {
		m.manufactureDate = nuevaFecha
	} else {
		fmt.Println("La fecha de fabricación no es válida")
	}
}

func (m *Medicine) SetShelfLife(nuevaShelfLife int) {
	if nuevaShelfLife > 0 {
		m.shelfLife = nuevaShelfLife
	} else {
		fmt.Println("Aun no se caduca pero pronto se caducara")
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

func NuevaTablet(name string, manufacturer string, mfgDate time.Time, shelfLife int, dose float64, prescription bool) Tablet {
	t := Tablet{
		Medicine:       NuevaMedicine(name, manufacturer, mfgDate, shelfLife),
		isPrescription: prescription,
	}
	t.SetDosePerTablet(dose)
	return t
}

func (t Tablet) GetDosePerTablet() float64 {
	return t.dosePerTablet
}

func (t Tablet) GetIsPrescription() bool {
	return t.isPrescription
}

func (t *Tablet) SetDosePerTablet(dose float64) {
	if dose > 0 {
		t.dosePerTablet = dose
	} else {
		fmt.Println("Error: La dosis por tableta debe ser mayor a 0.")
	}
}

func (t *Tablet) SetIsPrescription(prescription bool) {
	t.isPrescription = prescription
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
	




























		