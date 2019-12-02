use std::io;
use std::io::prelude::*;
use std::io::BufReader;
use std::fs::File;

fn get_fuel(mass: f64) -> f64 {
    ((mass / 3.0).floor() - 2.0).max(0.0)
}

fn get_total_fuel(mass: f64) -> f64 {
    fn get_fuel_rec(mass: f64, total: f64) -> f64 {
        let additional = get_fuel(mass);
        if additional > 0.0 {
            get_fuel_rec(additional, total + additional)
        } else {
            total
        }
    }
    get_fuel_rec(mass, 0.0)
}

fn calc(func: fn(f64) -> f64) -> io::Result<f64> {
    let f = File::open("input.txt")?;

    let total: f64 = BufReader::new(f)
        .lines()
        .map(|l| l.unwrap().parse::<f64>().unwrap())
        .map(func).sum();

    Ok(total)
}

fn main() {
    println!("Part one: {}", calc(get_fuel).unwrap());
    println!("Part two: {}", calc(get_total_fuel).unwrap());
}
