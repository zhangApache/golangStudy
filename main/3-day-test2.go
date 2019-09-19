package main

type Retangle struct {
	width, height float64
}

func area(r Retangle) float64 {
	return r.height * r.width
}
func main() {
	r1 := Retangle{10, 15}
	area(r1)
}