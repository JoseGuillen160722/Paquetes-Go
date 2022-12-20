package cliente

type Cliente struct {
	Id             int64
	NombreCompleto string
	Telefono       string
}

func HolaMundo() string {
	return "Hola mundo"
}
