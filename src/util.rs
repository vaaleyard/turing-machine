use std::{error::Error, path::Path};

use crate::model::Machine;

pub fn read_machine_from_file<P: AsRef<Path>>(path: P) -> Result<Machine, Box<dyn Error>> {
    // Open the file in read-only mode with buffer.
    let file = std::fs::File::open(path)?;
    let reader = std::io::BufReader::new(file);

    // Read the JSON contents of the file as an instance of `User`.
    let u = serde_json::from_reader(reader)?;

    // Return the `User`.
    Ok(u)
}
