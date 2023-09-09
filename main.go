package main

import (
	"io"
	"log"

	"github.com/AvoidMe/gpython/eval"
	pythontogo "github.com/AvoidMe/gpython/python_to_go"
)

func main() {
	log.SetOutput(io.Discard)
	instructions := pythontogo.LoadJson()
	eval.EvalInstructions(instructions)
}
