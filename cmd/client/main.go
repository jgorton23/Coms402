package main

import "github.com/MatthewBehnke/apis/internal/client/ui"

func main() {
	prog, err := ui.NewProgram()
	if err != nil {
		return
	}
	if err := prog.Start(); err != nil {
		return
	}
}
