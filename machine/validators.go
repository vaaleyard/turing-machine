package machine

import "log"

func (machine *Machine) ValidateChain(actualState string) bool {
	for _, finalStates := range machine.EstadosFinais {
		if finalStates == actualState {
			log.Println(actualState)
			return true
		}
	}
	return false
}
