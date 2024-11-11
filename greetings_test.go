package greetingsMudule

import (
	"regexp"
	"testing"
)

// testing.T es para reportar el resultado de la prueba
func TestHelloName(t *testing.T) {
	name := "Juan"
	// expresion regular, esto buscara una coincidencia exacta con el nombre
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Juan")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Juan") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "",error`, msg, err)
	}
}
