package main

import (
	"fmt"

	Cliente "./Cliente"
	Factura "./Factura"
	Producto "./Producto"
)

func main() {
	fmt.Println(Cliente.HolaMundo())
	fmt.Println(Producto.HolaProducto())
	fmt.Println(Factura.HolaFactura())
}
