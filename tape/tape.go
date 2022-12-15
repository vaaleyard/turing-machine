package tape

import (
	"fmt"
	"strings"

	. "github.com/vaaleyard/turing-machine/transition"
)

const TAPE_SIZE = 20

type Cell struct {
    Data string
}

type Tape struct {
    Tape []Cell
    Index uint
}

func (tape *Tape) WriteTape(alphabet string) {
    var chars []string = strings.Split(alphabet, "")

    for i := 0; i <= TAPE_SIZE; i++ {
        tape.Tape = append(tape.Tape, Cell{
            Data: chars[i],
        })
    }
}

func (tape *Tape) Move(transition Transition) {
    switch transition.Direcao {
    case "LEFT":
        tape.Index -= 1
    case "RIGHT":
        tape.Index += 1
    }
}

func (tape *Tape) Print() {
    for i := 0; i <= TAPE_SIZE; i++ {
        fmt.Print(tape.Tape[i].Data)
    }
    fmt.Println()
}
