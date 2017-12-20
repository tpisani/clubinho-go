package numbers

const (
	One = 1 // Exportado.
	two = 2 // Não exportado.
)

func Add(x, y int) int { // Exportado.
	return x + y
}

func sub(x, y int) int { // Não exportado.
	return x - y
}
