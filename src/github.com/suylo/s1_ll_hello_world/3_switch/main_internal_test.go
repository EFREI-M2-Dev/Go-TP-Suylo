package main

import "testing"

func Example_main() {
	main()
	// Output:
	// Hello, World!
}

func TestGreet(t *testing.T) {
	// Phase de prépa : on définit les valeurs attendues
	expectedGreeting := "Hello, World!"

	// Phase d'exécution : on appelle la fonction à tester
	greeting := greet("fr")

	// Phase décision : On vérifie si le résultat est conforme aux attentes
	if greeting != expectedGreeting {
		// Fail
		t.Errorf("Expected %q, but got %q", expectedGreeting, greeting)
	}
}
