use std::io;
use std::io::prelude::*;
use std::io::BufReader;
use std::fs::File;

fn load_intcode() -> io::Result<Vec<usize>> {
    let f = File::open("input.txt")?;

    let mut contents = String::new();
    BufReader::new(f).read_to_string(&mut contents).expect("couldn't read file");

    let intcode = contents
        .split(',')
        .map(|code| code.trim().parse::<usize>().unwrap())
        .collect();

    Ok(intcode)
}

fn execute(intcode: &mut Vec<usize>) {
    let mut pc = 0;
    loop {
        match intcode[pc] {
            1 | 2 => {
                let i_op1 = intcode[pc + 1];
                let i_op2 = intcode[pc + 2];
                let i_res = intcode[pc + 3];

                intcode[i_res] = match intcode[pc] {
                    1 => intcode[i_op1] + intcode[i_op2],
                    2 => intcode[i_op1] * intcode[i_op2],
                    _ => panic!("WAT"),
                };
                pc += 4;
            },
            99 => break,
            _ => panic!("Illegal opcode: {}", intcode[pc]),
        }
    }
}

fn main() -> io::Result<()> {
    let mut intcode = load_intcode().unwrap();

    // Restore 1202 program alarm
    intcode[1] = 12;
    intcode[2] = 2;

    execute(&mut intcode);

    println!("Position 0: {}", intcode[0]);

    Ok(())
}
