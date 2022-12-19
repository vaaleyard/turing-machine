package machine

import (
	. "github.com/vaaleyard/turing-machine/transition"
)

type Machine struct {
	Alfabeto       []string     `json:"alfabeto"`
	Simbolos       []string     `json:"simbolos"`
	Transicoes     []Transition `json:"transicoes"`
	MarcadorInicio string       `json:"marcadorInicio"`
	SimboloBranco  string       `json:"simboloBranco"`
	EstadoInicial  *string      `json:"estadoInicial"`
	EstadosFinais  []string     `json:"estadoFinais"`
	Estados        []string     `json:"estados"`
}
