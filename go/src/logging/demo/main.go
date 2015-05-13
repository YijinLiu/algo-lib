package main

import (
	"os"

	"logging"
)

func main() {
	for _, arg := range os.Args {
		logging.Print(arg)
		logging.Vlog(-1, arg)
		logging.Vlog(0, arg)
		logging.Vlog(1, arg)
	}
}
