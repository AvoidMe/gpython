package main

import (
	"io/ioutil"
	"log"

	"github.com/AvoidMe/gpython/eval"
	pythontogo "github.com/AvoidMe/gpython/python_to_go"
)

func main() {
	log.SetOutput(ioutil.Discard)
	instructions := pythontogo.LoadJson()
	eval.EvalInstructions(instructions)
}
