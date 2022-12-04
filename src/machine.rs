use std::error::Error;
use std::path::Path;

use crate::model::{Machine, Transition};
use crate::tape::Tape;

const BLANK_CELL_QUANTITY: usize = 50;

impl Machine {
    pub fn new_machine_from_file<P: AsRef<Path>>(path: P) -> Result<Machine, Box<dyn Error>> {
        // read-only
        let file = std::fs::File::open(path)?;
        let reader = std::io::BufReader::new(file);

        // read json as an instance of `Machine`
        let machine = serde_json::from_reader(reader)?;

        Ok(machine)
    }

    // Returns the alphabet with the init symbom and infinite right blank symbols
    fn construct_alphabet(&mut self, alphabet: String) -> String {
        let mut marcador_inicio: String = self.marcador_inicio.to_owned();
        marcador_inicio.push_str(alphabet.as_str());

        let quantity_of_blank_cells: usize = BLANK_CELL_QUANTITY - marcador_inicio.len();
        marcador_inicio.push_str(&self.simbolo_branco.repeat(quantity_of_blank_cells).as_str());

        marcador_inicio
    }

    fn create_tape(&mut self, alphabet: String) -> Tape {
        let full_alphabet_tape: String = self.construct_alphabet(alphabet);

        let mut tape: Tape = Tape::default();
        tape.write_tape(full_alphabet_tape);

        tape
    }

    fn find_transictions_by_actual_state(
        &mut self,
        actual_state: String,
        symbol: String,
    ) -> Vec<&Transition> {

        self.transicoes
            .iter()
            .filter(|transicao| {
                transicao.estado_origem == actual_state && transicao.simbolo == symbol
            })
            .collect::<_>()
    }

    fn is_done(&mut self, tape: &mut Tape, actual_state: &mut String) -> bool {
        return self
            .find_transictions_by_actual_state(
                actual_state.to_owned(),
                tape.tape[tape.index].data.to_owned(),
            )
            .is_empty();
    }

    fn step(&mut self, tape: &mut Tape, actual_state: &mut String) {
        let binding = self.find_transictions_by_actual_state(
            actual_state.to_owned(),
            tape.tape[tape.index].data.to_owned(),
        );

        let future_transition = match binding.first() {
            Some(transition) => transition,
            None => panic!("No more transitions"),
        };

        tape.tape[tape.index].data = future_transition.escreve.to_owned();
        *actual_state = future_transition.estado_destino.clone();

        tape.move_index(future_transition);
    }

    pub fn run_machine(&mut self, alphabet: String) {
        let mut actual_state = self.estado_inicial.to_owned();

        let mut tape: Tape = self.create_tape(alphabet);

        while !self.is_done(&mut tape, &mut actual_state) {
            self.step(&mut tape, &mut actual_state);
        }

        tape.print_tape();
    }
}
