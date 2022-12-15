package main

import (
	"github.com/vaaleyard/turing-machine/machine"
)


func main() {
    var machine machine.Machine = machine.NewMachineFromFile("examples/maquina01.json")
    var alphabet string = "000111"

    machine.Run(alphabet)
}
