package greetingsMudule

import (
	"errors"
	"fmt"
	"math/rand"
)

// hello devuelve un saludo para una persona especifica
// si esta en mayuscula sera publica para usarse fuera del paqueteß
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is required")
	}
	return fmt.Sprintf(randomFormat(), name), nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil

}

func randomFormat() string {
	formats := []string{
		"Hola!, bienvenido %s!",
		"buena mi perro!, el %s!",
		"que se narra la mojarra!, %s!",
	}
	return formats[rand.Intn(len(formats))]
}
