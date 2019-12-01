use std::io;
use std::io::prelude::*;
use std::io::BufReader;
use std::fs::File;

fn get_fuel(mass: f64) -> f64 {
    let mut total = 0.0;
    let mut balance = mass;
    loop {
        let additional = ((balance / 3.0).floor() - 2.0).max(0.0);
        if additional > 0.0 {
            total += additional;
            balance = additional
        } else {
            // We've accounted for all the additional fuel
            break
        }
    }
    total
}

fn main() -> io::Result<()> {
    let f = File::open("input.txt")?;
    let reader = BufReader::new(f);

    let mut total = 0.0;
    for line in reader.lines() {
        let value: f64 = line?.parse().unwrap();
        total += get_fuel(value);
    }

    println!("Fuel requirement: {}", total);

    Ok(())
}
