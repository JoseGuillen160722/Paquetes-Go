package producto

type Producto struct {
	Id             int64
	NombreProducto string
	Precio         float64
}

func HolaProducto() string {
	return "Hola producto"
}
