mod machine;
mod model;
mod tape;
mod util;

use std::io;

use crate::model::Machine;

fn main() -> io::Result<()> {
    let mut machine: Machine = util::read_machine_from_file("maquina.json").unwrap();

    // let mut input = String::new();
    // let stdin = io::stdin();
    // stdin.read_line(&mut input)?;
    // let alphabet: String = input.trim().to_string();
    let alphabet: String = "0011".to_string();

    machine.process_machine(alphabet);

    Ok(())
}
