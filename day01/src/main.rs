use std::io;
use std::io::prelude::*;
use std::io::BufReader;
use std::fs::File;

fn main() -> io::Result<()> {
    let f = File::open("input.txt")?;
    let reader = BufReader::new(f);

    let mut total: i32 = 0;
    for line in reader.lines() {
        total += line?.parse::<i32>().unwrap()
    }
    println!("{}", total);

    Ok(())
}
