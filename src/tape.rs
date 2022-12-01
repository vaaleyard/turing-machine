const TAPE_SIZE: usize = 20;

#[derive(Default, Debug)]
pub struct Cell {
    pub data: String,
}

#[derive(Default)]
pub struct Tape {
    pub tape: Vec<Cell>,
    pub index: usize,
}

impl Tape {
    // Set the alphabet to the Tape struct
    pub fn write_tape(&mut self, alphabet: String) {
        let char_vec: Vec<char> = alphabet.chars().collect();

        for i in 0..TAPE_SIZE {
            let mut cell: Vec<Cell> = vec![Cell {
                data: char_vec[i].to_string(),
            }];
            self.tape.append(&mut cell);
        }
    }
}
