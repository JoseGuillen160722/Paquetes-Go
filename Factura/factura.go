package factura

type Factura struct {
	Id     int64
	Pais   string
	Ciudad string
	Total  float64
	client cliente.Cliente
}

func HolaFactura() string {
	return "Hola Factura"
}
