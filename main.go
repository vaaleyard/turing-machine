package main

import (
	"github.com/vaaleyard/turing-machine/machine"
	"github.com/vaaleyard/turing-machine/ui"
)


func main() {
    var machine machine.Machine = machine.NewMachineFromFile("examples/maquina01.json")
    var alphabet string = "01"

    machine.Run(alphabet)

    ui.Ui()
}
