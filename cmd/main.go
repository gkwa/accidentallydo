package main

import (
	"os"

	"github.com/taylormonacelli/accidentallydo"
)

func main() {
	code := accidentallydo.Execute()
	os.Exit(code)
}
