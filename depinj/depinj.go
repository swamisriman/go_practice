package depinj

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, personName string) {
	fmt.Fprintf(writer, "Hello, "+personName)
}
