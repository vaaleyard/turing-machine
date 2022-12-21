package machine

func (machine *Machine) ValidateChain(actualState string) bool {
	for _, finalStates := range machine.EstadosFinais {
		if finalStates == actualState {
			return true
		}
	}
	return false
}
