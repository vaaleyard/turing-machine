package machine

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/vaaleyard/turing-machine/tape"
	. "github.com/vaaleyard/turing-machine/transition"
)

const BLANK_CELL_QUANTITY = 50

func NewMachineFromFile(path string) Machine {
    jsonFile, err := os.Open(path)
    if err != nil {
        log.Panicln("Failed to open current file")
    }

    byteJsonFileValue, _ := ioutil.ReadAll(jsonFile)

    var machine Machine
    json.Unmarshal(byteJsonFileValue, &machine)
    defer jsonFile.Close()

    return machine
}

func (machine *Machine) constructAlphabet(alphabet string) string {
    var marcadorInicio string = machine.MarcadorInicio
    marcadorInicio += alphabet

    var quantityOfBlankCells int = BLANK_CELL_QUANTITY - len(marcadorInicio)
    var fullAlphabetWithBlankSymbols string = marcadorInicio +
            strings.Repeat(machine.SimboloBranco, quantityOfBlankCells)

    return fullAlphabetWithBlankSymbols
}

func (machine *Machine) createTape(alphabet string) *tape.Tape {
    var full_alphabet string = machine.constructAlphabet(alphabet)

    var tape *tape.Tape = new(tape.Tape)
    tape.WriteTape(full_alphabet)

    return tape
}

func (machine *Machine) findTransitionsByActualState(actualState string, symbol string) []Transition {
    var matchedTransitions []Transition

    for _, transition := range machine.Transicoes {
        if transition.EstadoOrigem == actualState &&
            transition.LeSimbolo == symbol {
            matchedTransitions = append(matchedTransitions, transition)     
        }
    }

    return matchedTransitions
}

func (machine *Machine) isDone(tape *tape.Tape, actualState *string) bool {
    var transitions []Transition = machine.findTransitionsByActualState(*actualState,
                                                tape.Tape[tape.Index].Data)

    if len(transitions) == 0 {
        return true
    } else {
        return false
    }
}

func (machine *Machine) step(tape *tape.Tape, actualState *string) bool {
    var binding []Transition = machine.findTransitionsByActualState(*actualState,
                                                tape.Tape[tape.Index].Data)

    var futureTransition Transition = binding[0]
    tape.Tape[tape.Index].Data = futureTransition.Escreve
    *actualState = futureTransition.EstadoDestino

    tape.Move(futureTransition)
    return true
}

func (machine *Machine) Run(alphabet string) {
    var actualState *string = machine.EstadoInicial
    tape := machine.createTape(alphabet)

    for !machine.isDone(tape, actualState) {
        machine.step(tape, actualState)
    }

    tape.Print()
}
