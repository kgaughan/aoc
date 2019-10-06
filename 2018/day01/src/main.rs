use std::collections::HashSet;
use std::io;
use std::io::prelude::*;
use std::io::BufReader;
use std::fs::File;
use std::vec;

fn main() -> io::Result<()> {
    let f = File::open("input.txt")?;
    let reader = BufReader::new(f);

    let mut total: i32 = 0;
    let mut values: vec::Vec<i32> = vec![];
    for line in reader.lines() {
        let value: i32 = line?.parse().unwrap();
        values.push(value);
        total += value;
    }

    let mut first_repeated: i32 = 0;
    let mut reached: HashSet<i32> = HashSet::new();
    let mut found = false;
    let mut freq = 0;
    while !found {
        for value in &values {
            reached.insert(freq);
            freq += value;
            if reached.contains(&freq) {
                found = true;
                first_repeated = freq;
                break;
            }
        }
    }

    println!("Final frequency: {}", total);
    if found {
        println!("First repeated: {}", first_repeated);
    } else {
        println!("No repetition found")
    }

    Ok(())
}
