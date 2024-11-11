package greetingsMudule

import (
	"fmt"
)

// hello devuelve un saludo para una persona especifica
// si esta en mayuscula sera publica para usarse fuera del paqueteß
func Hello(name string) string {
	return fmt.Sprintf("Hola!, bienvenido %s!", name)
}
