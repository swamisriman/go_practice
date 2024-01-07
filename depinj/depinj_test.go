package depinj

import (
	"bytes"
	// "os"
	// "net/http"
	"testing"
)

func TestGreet(t *testing.T) {
	buff := bytes.Buffer{}
	Greet(&buff, "Sriman")

	expected := "Hello, Sriman"
	actual := buff.String()

	if expected != actual {
		t.Errorf("Expected: %q but Actual: %q", expected, actual)
	}

	// Greet(os.Stdout, "Sriman")

	// http.ListenAndServe(":5001", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	Greet(w, "world")
	// }))
}
