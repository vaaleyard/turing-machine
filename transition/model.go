package transition

type Transition struct {
    LeSimbolo string `json:"leSimbolo"`
    Escreve string `json:"escreve"`
    Direcao string `json:"direcao"`
    EstadoOrigem string `json:"estadoOrigem"`
    EstadoDestino string `json:"estadoDestino"`
}


