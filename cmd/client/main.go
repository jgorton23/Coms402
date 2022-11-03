package main

import "github.com/MatthewBehnke/exampleGoApi/cmd/client/ui"

func main() {
	prog, err := ui.NewProgram()
	if err != nil {
		return
	}
	if err := prog.Start(); err != nil {
		return
	}
}
