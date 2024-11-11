package greetingsMudule

import (
	"fmt"
)

// hello devuelve un saludo para una persona especifica
func hello(name string) string {
	return fmt.Sprintf("Hola!, bienvenido %s!", name)
}
