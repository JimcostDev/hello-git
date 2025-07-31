package main

import (
	"bytes"
	"os"
	"testing"
)

// TestGetGreeting prueba la función GetGreeting
func TestGetGreeting(t *testing.T) {
	expected := "Hola Mundo en Go!"
	result := GetGreeting()

	if result != expected {
		t.Errorf("GetGreeting() = %q, esperado %q", result, expected)
	}
}

// TestPrintGreeting prueba la función PrintGreeting
func TestPrintGreeting(t *testing.T) {
	// Capturamos la salida estándar
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Ejecutamos la función
	PrintGreeting()

	// Restauramos la salida estándar
	w.Close()
	os.Stdout = old

	// Leemos lo que se imprimió
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	expected := "Hola Mundo en Go!\n"
	if output != expected {
		t.Errorf("PrintGreeting() imprimió %q, esperado %q", output, expected)
	}
}

// BenchmarkGetGreeting benchmark para la función GetGreeting
func BenchmarkGetGreeting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetGreeting()
	}
}

// TestGetGreetingNotEmpty verifica que el saludo no esté vacío
func TestGetGreetingNotEmpty(t *testing.T) {
	result := GetGreeting()

	if result == "" {
		t.Error("GetGreeting() no debería retornar una cadena vacía")
	}
}

// TestGetGreetingContainsGo verifica que el saludo contenga "Go"
func TestGetGreetingContainsGo(t *testing.T) {
	result := GetGreeting()

	if !containsSubstring(result, "Go") {
		t.Errorf("GetGreeting() = %q, debería contener 'Go'", result)
	}
}

// Función auxiliar para verificar si una cadena contiene una subcadena
func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
