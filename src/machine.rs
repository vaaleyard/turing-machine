use crate::model::{Machine, Transition, Move};
use crate::tape::Tape;

const BLANK_CELL_QUANTITY: usize = 50;

impl Machine {
    // Returns the alphabet with the init symbom and infinite right blank symbols
    pub fn construct_alphabet(&mut self, alphabet: String) -> String {
        let mut owned_marcador_inicio: String = self.marcador_inicio.to_owned();
        owned_marcador_inicio.push_str(alphabet.as_str());

        let quantity_of_blank_cells: usize = BLANK_CELL_QUANTITY - owned_marcador_inicio.len();
        owned_marcador_inicio.push_str(&self.simbolo_branco.repeat(quantity_of_blank_cells).as_str());

        owned_marcador_inicio
    }

    pub fn create_tape(&mut self, alphabet: String) -> Tape {
        let full_alphabet_tape: String = self.construct_alphabet(alphabet);

        let mut tape: Tape = Tape::default();
        tape.write_tape(full_alphabet_tape);

        tape
    }

    fn find_transictions_by_actual_state(&mut self, actual_state: String, symbol: String) -> Vec<&Transition> {
        self.transicoes.iter().filter(|transicao| transicao.estado_origem == actual_state && transicao.simbolo == symbol).collect::<Vec<_>>()
    }

    fn is_done(&mut self, tape: &mut Tape, actual_state: &mut String) -> bool {
        for symbol in &tape.tape {
            return self.find_transictions_by_actual_state(actual_state.to_owned(), symbol.data.to_owned()).is_empty()
        }
        true
    }

    fn move_index(&mut self, tape: &mut Tape, transition: Transition) {
        if transition.direcao == Move::LEFT {
            tape.index -= 1;
        } else if transition.direcao == Move::RIGHT {
            tape.index += 1;
        }
    }

    pub fn step(&mut self, tape: &mut Tape, actual_state: &mut String) -> String {
        for symbol in &tape.tape {
            let transitions = self.find_transictions_by_actual_state(actual_state.to_owned(), symbol.data.to_owned());

            let next_transition = match transitions.first() {
                Some(transition) => Some(transition),
                None => None,
            };
            *actual_state = next_transition.unwrap().estado_destino.clone();
            println!("simbolo: {:?} actual: {:?} // next: {:?}", symbol.data, actual_state, next_transition);

            break;
        }
        return actual_state.to_owned();
    }

    pub fn process_machine(&mut self, alphabet: String) {
        let mut actual_state = self.estado_inicial.to_owned();

        let mut tape: Tape = self.create_tape(alphabet);

        while !self.is_done(&mut tape, &mut actual_state) {
            actual_state = self.step(&mut tape, &mut actual_state);
            break;
        }
    }

}
