package main

import (
	"os"

	"/logging"
)

func main() {
	for _, arg := range os.Args {
		logging.Print(arg)
	}
}
