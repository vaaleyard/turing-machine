mod machine;
mod model;
mod tape;

use std::io::Result;

use crate::model::Machine;

fn main() -> Result<()> {
    let mut machine: Machine = Machine::new_machine_from_file("maquina.json").unwrap();
    let alphabet: String = "0011".to_string();

    machine.run_machine(alphabet);

    Ok(())
}
